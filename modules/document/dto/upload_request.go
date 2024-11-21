package dto

type PresignedURLRequestBody struct {
	Name   string `json:"name" validate:"required" example:"example.pdf"`
	Action string `json:"action" validate:"required" example:"upload"`
}

type UploadDocumentRequestBody struct {
	ObjectKey string       `json:"key" validate:"required" example:"example.pdf"`
	MetaData  MetaDataBody `json:"metadata" validate:"required"`
}

type MetaDataBody struct {
	Name      string `json:"name"`
	Size      int64  `json:"size"`
	MimeType  string `json:"mime_type"`
	Extension string `json:"extension"`
	Path      string `json:"path"`
}
