package user

import "github.com/mixarchitecture/arch/shared/i18n"

type Errors interface {
	NotFound(email string) *i18n.I18nError
	AlreadyExists(email string) *i18n.I18nError
	Failed(operation string) *i18n.I18nError
	InvalidPassword() *i18n.I18nError
}

type userErrors struct{}

func newUserErrors() Errors {
	return &userErrors{}
}

func (e *userErrors) NotFound(email string) *i18n.I18nError {
	return i18n.NewError(I18nMessages.NotFound, i18n.P{"Email": email})
}

func (e *userErrors) AlreadyExists(email string) *i18n.I18nError {
	return i18n.NewError(I18nMessages.AlreadyExists, i18n.P{"Email": email})
}

func (e *userErrors) Failed(operation string) *i18n.I18nError {
	return i18n.NewError(I18nMessages.Failed, i18n.P{"Operation": operation})
}

func (e *userErrors) InvalidPassword() *i18n.I18nError {
	return i18n.NewError(I18nMessages.InvalidPassword)
}
