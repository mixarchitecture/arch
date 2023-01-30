package example

import "errors"

type FactoryConfig struct {
	MinFieldLength   int
	MaxFieldLength   int
	MinContentLength int
	MaxContentLength int
}

func (f FactoryConfig) Validate() error {
	if f.MinFieldLength < 0 {
		return errors.New("min field length must be greater than or equal to 0")
	}
	if f.MaxFieldLength < 0 {
		return errors.New("max field length must be greater than or equal to 0")
	}
	if f.MinFieldLength > f.MaxFieldLength {
		return errors.New("min field length must be less than or equal to max field length")
	}
	if f.MinContentLength < 0 {
		return errors.New("min content length must be greater than or equal to 0")
	}
	if f.MaxContentLength < 0 {
		return errors.New("max content length must be greater than or equal to 0")
	}
	if f.MinContentLength > f.MaxContentLength {
		return errors.New("min content length must be less than or equal to max content length")
	}
	return nil
}
