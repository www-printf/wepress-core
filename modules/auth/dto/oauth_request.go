package dto

type OauthCallbackRequestBody struct {
	Provider string `json:"provider" validate:"required" example:"provider"`
	Code     string `json:"code" validate:"required" example:"code"`
	State    string `json:"state" validate:"required" example:"state"`
	Error    string `json:"error"`
}

type OauthCallBackTransfer struct {
	Code     string
	State    string
	Verifier string
}
