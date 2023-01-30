package dtos

import "github.com/mixarchitecture/arch/example/src/domain/example"

type CreateExampleResponse struct {
	Field   string `json:"field"`
	Content string `json:"content"`
}

type ListExampleResponse struct {
	Offset int               `json:"offset"`
	Limit  int               `json:"limit"`
	Total  int               `json:"total"`
	Items  []example.Example `json:"items"`
}

type GetExampleResponse struct {
	UUID    string `json:"uuid"`
	Field   string `json:"field"`
	Content string `json:"content"`
}

type UpdateExampleResponse struct {
	Field   string `json:"field"`
	Content string `json:"content"`
}
