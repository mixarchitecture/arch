package http

type successMessages struct {
	Login    string
	Logout   string
	Register string
	Extend   string
}

type errorMessages struct {
	Login      string
	Logout     string
	Register   string
	Extend     string
	Unexpected string
}

type messages struct {
	Success successMessages
	Error   errorMessages
}

var Messages = messages{
	Success: successMessages{
		Login:    "success_login",
		Logout:   "success_logout",
		Register: "success_register",
		Extend:   "success_extend",
	},
	Error: errorMessages{
		Login:      "error_login",
		Logout:     "error_logout",
		Register:   "error_register",
		Extend:     "error_extend",
		Unexpected: "error_unexpected",
	},
}
