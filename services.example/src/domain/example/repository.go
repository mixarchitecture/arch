package example

import (
	"context"

	"github.com/mixarchitecture/arch/shared/i18n"
)

type Repository interface {
	Get(ctx context.Context, field string) (*Example, *i18n.I18nError)
	List(ctx context.Context, limit int, offset int) ([]*Example, int, *i18n.I18nError)

	Create(ctx context.Context, example *Example) (*Example, *i18n.I18nError)
	Update(ctx context.Context, example *Example) (*Example, *i18n.I18nError)
}
