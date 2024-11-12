package dto

type UserResponseBody struct {
	ID        string `json:"id"`
	Fullname  string `json:"fullname"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	PubKey    string `json:"pubkey"`
}
