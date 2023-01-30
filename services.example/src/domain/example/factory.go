package example

import (
	"github.com/mixarchitecture/arch/shared/i18n"

	"github.com/pkg/errors"
)

type Factory struct {
	fc FactoryConfig
}

func NewFactory(fc FactoryConfig) (Factory, error) {
	if err := fc.Validate(); err != nil {
		return Factory{}, errors.Wrap(err, "invalid factory config")
	}
	return Factory{fc: fc}, nil
}

func MustNewFactory(fc FactoryConfig) Factory {
	f, err := NewFactory(fc)
	if err != nil {
		panic(err)
	}
	return f
}

func (f Factory) Config() FactoryConfig {
	return f.fc
}

func (f Factory) IsZero() bool {
	return f == Factory{}
}

func (f Factory) NewExample(field string, content string) (*Example, *i18n.I18nError) {
	if err := f.validateField(field); err != nil {
		return nil, err
	}
	if err := f.validateContent(content); err != nil {
		return nil, err
	}
	return &Example{
		Field:   field,
		Content: content,
	}, nil
}

func (f Factory) Unmarshal(uuid string, field string, content string) (*Example, *i18n.I18nError) {
	if err := f.validateField(field); err != nil {
		return nil, err
	}
	if err := f.validateContent(content); err != nil {
		return nil, err
	}
	return &Example{
		UUID:    uuid,
		Field:   field,
		Content: content,
	}, nil
}

func (f Factory) Validate(e *Example) *i18n.I18nError {
	if err := f.validateField(e.Field); err != nil {
		return err
	}
	if err := f.validateContent(e.Content); err != nil {
		return err
	}
	return nil
}

func (f Factory) validateField(field string) *i18n.I18nError {
	if len(field) < f.fc.MinFieldLength {
		return i18n.NewError(I18nMessages.Field_Too_Short, i18n.P{"Param": f.fc.MinFieldLength})
	}
	if len(field) > f.fc.MaxFieldLength {
		return i18n.NewError(I18nMessages.Field_Too_long, i18n.P{"Param": f.fc.MaxFieldLength})
	}
	return nil
}

func (f Factory) validateContent(content string) *i18n.I18nError {
	if len(content) < f.fc.MinContentLength {
		return i18n.NewError(I18nMessages.Content_Too_Short, i18n.P{"Param": f.fc.MinContentLength})
	}
	if len(content) > f.fc.MaxContentLength {
		return i18n.NewError(I18nMessages.Content_Too_long, i18n.P{"Param": f.fc.MaxContentLength})
	}
	return nil
}

func (f Factory) NewNotFoundError(key string) *i18n.I18nError {
	return i18n.NewError(I18nMessages.Not_Found, i18n.P{"Param": key})
}
