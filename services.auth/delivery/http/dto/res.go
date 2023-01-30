package dto

type RegisterResponse struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
