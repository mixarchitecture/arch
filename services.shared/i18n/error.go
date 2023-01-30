package i18n

type I18nError struct {
	Key    string
	Params interface{}
}

type P map[string]interface{}

func (e *I18nError) Error() string {
	return e.Key
}

func (e *I18nError) IsErr() bool {
	return e.Key != ""
}

func NewError(key string, params ...P) *I18nError {
	p := &P{}
	if len(params) > 0 {
		p = &params[0]
	}
	return &I18nError{Key: key, Params: p}
}
