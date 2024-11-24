package dto

type OauthResponseBody struct {
	URL string `json:"url"`
}

type OauthUserResponse struct {
	Email    string `json:"email"`
}

type GithubEmail struct {
	Email    string `json:"email"`
}