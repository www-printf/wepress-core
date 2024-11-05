package dto

type AuthResponse struct {
	Token string `json:"token" validate:"require,jwt"`
	Type  string `json:"type"`
}
