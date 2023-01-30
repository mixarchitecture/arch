package user

import (
	"github.com/google/uuid"
	"github.com/mixarchitecture/arch/auth/domain/user"
	"github.com/mixarchitecture/arch/shared/i18n"
	"golang.org/x/net/context"
)

func (r *repo) Create(ctx context.Context, email string, password []byte) (*user.User, *i18n.I18nError) {
	user := r.userFactory.NewUser(email, password)
	user.UUID = uuid.New().String()
	r.users[user.UUID] = *user
	return user, nil
}

func (r *repo) Update(ctx context.Context, user *user.User) (*user.User, *i18n.I18nError) {
	r.users[user.UUID] = *user
	return user, nil
}

func (r *repo) GetByEmail(ctx context.Context, email string) (*user.User, *i18n.I18nError) {
	for _, u := range r.users {
		if u.Email == email {
			return &u, nil
		}
	}
	return nil, r.userFactory.Errors.NotFound(email)
}
