package dto

type PresignedURLResponseBody struct {
	URL       string `json:"url" validate:"required"`
	ObjectKey string `json:"object_key" validate:"required"`
}
