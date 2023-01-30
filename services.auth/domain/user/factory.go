package user

import (
	"time"

	"github.com/mixarchitecture/arch/shared/formats"
	"github.com/mixarchitecture/arch/shared/i18n"
)

type Factory struct {
	Errors Errors
}

func NewFactory() Factory {
	return Factory{
		Errors: newUserErrors(),
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}

func (f Factory) NewUser(email string, password []byte) *User {
	t := time.Now().Format(formats.ISO8600)
	return &User{
		UUID:      "",
		Email:     email,
		Password:  password,
		Roles:     []string{"user"},
		IsActive:  true,
		CreatedAt: t,
		UpdatedAt: t,
	}
}

func (f Factory) Unmarshal(uuid string, email string, isActive bool) *User {
	return &User{
		UUID:     uuid,
		Email:    email,
		IsActive: isActive,
	}
}

func (f Factory) Validate(u *User) *i18n.I18nError {
	if err := f.validateEmail(u.Email); err != nil {
		return err
	}
	return nil
}

func (f Factory) validateEmail(email string) *i18n.I18nError {
	if email == "" {
		return i18n.NewError("error_user_email_empty")
	}
	return nil
}
