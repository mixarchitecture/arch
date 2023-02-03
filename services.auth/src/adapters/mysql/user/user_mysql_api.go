package user

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mixarchitecture/arch/auth/src/adapters/mysql/user/entity"
	"github.com/mixarchitecture/arch/auth/src/domain/user"
	"github.com/mixarchitecture/arch/shared/formats"
	"github.com/mixarchitecture/arch/shared/i18n"
	sqb_go "gitlab.com/ssibrahimbas/sqb.go"
)

func (r *repo) Create(ctx context.Context, email string, password []byte) (*user.User, *i18n.I18nError) {
	user := r.userFactory.NewUser(email, password)
	e := r.checkExist(ctx, user.Email, false)
	if e != nil {
		return nil, r.userFactory.Errors.AlreadyExists(user.Email)
	}
	user.UUID = uuid.New().String()
	query := sqb_go.QB.Table(entity.Fields.Table).Insert(r.mapper.ToEntityMap(user))
	_, err := r.db.ExecContext(ctx, query)
	if err != nil {
		return nil, r.userFactory.Errors.Failed("create")
	}
	return user, nil
}

func (r *repo) Update(ctx context.Context, user *user.User) (*user.User, *i18n.I18nError) {
	e := r.checkExist(ctx, user.Email, true)
	if e != nil {
		return nil, e
	}
	query := sqb_go.QB.Table(entity.Fields.Table).Where(entity.Fields.UUID, "=", user.UUID).Update(&sqb_go.M{
		entity.Fields.IsActive:  user.IsActive,
		entity.Fields.UpdatedAt: time.Now().Format(formats.DateYYYYMMDDHHMMSS),
	})
	_, err := r.db.ExecContext(ctx, query)
	if err != nil {
		return nil, r.userFactory.Errors.Failed("update")
	}
	return user, nil
}

func (r *repo) GetByEmail(ctx context.Context, email string) (*user.User, *i18n.I18nError) {
	e := &entity.MySQLUser{}
	query := sqb_go.QB.Table(entity.Fields.Table).Where(entity.Fields.Email, "=", email).Get()
	err := r.db.GetContext(ctx, e, query)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, r.userFactory.Errors.NotFound(email)
	}
	if err != nil {
		return nil, r.userFactory.Errors.Failed("get")
	}
	return e.ToUser(), nil
}

func (r *repo) checkExist(ctx context.Context, email string, throwNotFound bool) *i18n.I18nError {
	e := &entity.MySQLUser{}
	query := sqb_go.QB.Table(entity.Fields.Table).Where(entity.Fields.Email, "=", email).Get()
	err := r.db.GetContext(ctx, e, query)
	if errors.Is(err, sql.ErrNoRows) && throwNotFound {
		return r.userFactory.Errors.NotFound(email)
	}
	if err != nil {
		return r.userFactory.Errors.Failed("checkExist")
	}
	return nil
}
