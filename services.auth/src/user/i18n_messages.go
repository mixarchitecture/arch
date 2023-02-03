package user

type messages struct {
	EmailEmpty      string
	NotFound        string
	AlreadyExists   string
	Failed          string
	InvalidPassword string
}

var I18nMessages = messages{
	EmailEmpty:      "error_user_email_empty",
	NotFound:        "error_user_not_found",
	Failed:          "error_user_failed",
	AlreadyExists:   "error_user_already_exists",
	InvalidPassword: "error_user_invalid_password",
}
