package dto

type AuthResponse struct {
	Token string `json:"token"`
	Type  string `json:"type"`
}
