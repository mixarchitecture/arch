package http

type successMessages struct {
	GetExample    string
	ListExample   string
	CreateExample string
	UpdateExample string
}

type errorMessages struct {
	GetExample    string
	ListExample   string
	CreateExample string
	UpdateExample string
	Unexpected    string
}

type messages struct {
	Success successMessages
	Error   errorMessages
}

var Messages = messages{
	Success: successMessages{
		GetExample:    "success_example_get",
		ListExample:   "success_example_list",
		CreateExample: "success_example_create",
		UpdateExample: "success_example_update",
	},
	Error: errorMessages{
		GetExample:    "error_example_get",
		ListExample:   "error_example_list",
		CreateExample: "error_example_create",
		UpdateExample: "error_example_update",
		Unexpected:    "error_unexpected",
	},
}
