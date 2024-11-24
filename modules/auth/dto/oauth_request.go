package dto

type OauthCallbackRequestBody struct {
	Provider string `json:"provider"`
	Code     string `json:"code" validate:"required" example:"code"`
	State    string `json:"state" validate:"required" example:"state"`
	Error    string `json:"error,omitempty"`
}

type OauthCallBackTransfer struct {
	Code     string
	Verifier string
}
