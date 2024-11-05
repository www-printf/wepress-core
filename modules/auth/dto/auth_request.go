package dto

type LoginRequestBody struct {
	Email    string `json:"email" validate:"required,email" example:"example@email.com"`
	Password string `json:"password" validate:"required,min=8,max=30" example:"password"`
}