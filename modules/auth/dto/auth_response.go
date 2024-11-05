package dto

type Token string

type AuthResponse struct {
	Token Token  `json:"token" validate:"require,jwt"`
	Type  string `json:"type"`
}
