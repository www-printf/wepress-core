package dto

type AuthResponseBody struct {
	Token string `json:"token" validate:"require,jwt"`
	Type  string `json:"type"`
}
