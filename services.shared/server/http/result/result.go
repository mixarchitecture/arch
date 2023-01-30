package result

import "github.com/gofiber/fiber/v2"

type Result struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type DetailResult struct {
	Result
	Detail any `json:"detail"`
}

func (r *Result) Error() string {
	return r.Message
}

func (r *DetailResult) Error() string {
	return r.Message
}

func Success(m string, c ...int) *Result {
	code := fiber.StatusOK
	if len(c) > 0 && c[0] != 0 {
		code = c[0]
	}
	return &Result{
		Message: m,
		Status:  code,
	}
}

func Error(m string, c ...int) *Result {
	code := fiber.StatusBadRequest
	if len(c) > 0 && c[0] != 0 {
		code = c[0]
	}
	return &Result{
		Message: m,
		Status:  code,
	}
}

func SuccessDetail(m string, d any, c ...int) *DetailResult {
	code := fiber.StatusOK
	if len(c) > 0 && c[0] != 0 {
		code = c[0]
	}
	return &DetailResult{
		Detail: d,
		Result: Result{Message: m, Status: code},
	}
}

func ErrorDetail(m string, d any, c ...int) *DetailResult {
	code := fiber.StatusBadRequest
	if len(c) > 0 && c[0] != 0 {
		code = c[0]
	}
	return &DetailResult{
		Detail: d,
		Result: Result{Message: m, Status: code},
	}
}
