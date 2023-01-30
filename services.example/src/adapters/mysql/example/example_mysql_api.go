package mysql_example

import (
	"context"
	"database/sql"

	"github.com/mixarchitecture/arch/shared/i18n"

	"github.com/mixarchitecture/arch/example/src/adapters/mysql/example/entity"
	"github.com/mixarchitecture/arch/example/src/domain/example"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	sqb_go "gitlab.com/ssibrahimbas/sqb.go"
)

func (r *exampleRepo) Get(ctx context.Context, field string) (*example.Example, *i18n.I18nError) {
	e := entity.MySQLExample{}
	query := sqb_go.QB.Table(entity.Fields.Table).Where(entity.Fields.Field, "=", field).Get()
	err := r.db.GetContext(ctx, &e, query)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, r.exampleFactory.NewNotFoundError(field)
	}
	if err != nil {
		return nil, i18n.NewError(example.I18nMessages.Get_Failed)
	}
	return r.mapper.ToExample(&e)
}

func (r *exampleRepo) Create(ctx context.Context, e *example.Example) (*example.Example, *i18n.I18nError) {
	if err := r.exampleFactory.Validate(e); err != nil {
		return nil, err
	}
	query := sqb_go.QB.Table(entity.Fields.Table).Insert(&sqb_go.M{
		entity.Fields.UUID:    uuid.New(),
		entity.Fields.Field:   e.Field,
		entity.Fields.Content: e.Content,
	})
	_, err := r.db.ExecContext(ctx, query)
	if err != nil {
		return nil, i18n.NewError(example.I18nMessages.Create_Failed)
	}
	return e, nil
}

func (r *exampleRepo) Update(ctx context.Context, e *example.Example) (*example.Example, *i18n.I18nError) {
	query := sqb_go.QB.Table(entity.Fields.Table).Where(entity.Fields.Field, "=", e.Field).Update(&sqb_go.M{
		entity.Fields.Content: e.Content,
	})
	_, err := r.db.ExecContext(ctx, query)
	if err != nil {
		return nil, i18n.NewError(example.I18nMessages.Update_Failed)
	}
	return e, nil
}

func (r *exampleRepo) List(ctx context.Context, limit int, offset int) ([]*example.Example, int, *i18n.I18nError) {
	var examples []*entity.MySQLExample
	query := sqb_go.QB.Table(entity.Fields.Table).Limit(limit).Offset(offset).Get()
	err := r.db.SelectContext(ctx, &examples, query)
	if err != nil {
		return nil, 0, i18n.NewError(example.I18nMessages.List_Failed)
	}
	count, error := r.Count(ctx)
	if error != nil {
		return nil, 0, error
	}
	res, error := r.mapper.ToExamples(examples)
	return res, count, error
}

func (r *exampleRepo) Count(ctx context.Context) (int, *i18n.I18nError) {
	var count int
	query := sqb_go.QB.Table(entity.Fields.Table).Count(entity.Fields.UUID, "count").Get()
	err := r.db.GetContext(ctx, &count, query)
	if err != nil {
		return 0, i18n.NewError(example.I18nMessages.Count_Failed)
	}
	return count, nil
}
