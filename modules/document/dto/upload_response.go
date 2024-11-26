package dto

type RequestUploadResponseBody struct {
	URL       string            `json:"url" validate:"required"`
	ObjectKey string            `json:"object_key" validate:"required"`
	Fields    map[string]string `json:"fields"`
}

type UploadResponseBody struct {
	ID       string       `json:"id" validate:"required"`
	MetaData MetaDataBody `json:"metadata" validate:"required"`
}
