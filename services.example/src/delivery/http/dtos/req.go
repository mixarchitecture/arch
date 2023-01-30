package dtos

type CreateExampleRequest struct {
	Field   string `json:"field" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type ListExampleRequest struct {
	Offset *int `query:"offset"  validate:"omitempty,gt=0"`
	Limit  *int `query:"limit"  validate:"omitempty,gt=0"`
}

type GetExampleRequest struct {
	Field string `param:"field" validate:"required"`
}

type UpdateExampleRequest struct {
	Field   string `param:"field" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func (r *ListExampleRequest) Default() {
	if r.Offset == nil {
		r.Offset = new(int)
		*r.Offset = 0
	}
	if r.Limit == nil {
		r.Limit = new(int)
		*r.Limit = 10
	}
}
