package user

import (
	"context"

	"github.com/mixarchitecture/arch/shared/i18n"
)

type Repository interface {
	GetByEmail(ctx context.Context, email string) (*User, *i18n.I18nError)
	Create(ctx context.Context, email string, password []byte) (*User, *i18n.I18nError)
	Update(ctx context.Context, user *User) (*User, *i18n.I18nError)
}
