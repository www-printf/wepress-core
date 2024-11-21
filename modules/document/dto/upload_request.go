package dto

type UploadRequestBody struct {
	RequestSize int64 `json:"size" validate:"required" example:"10485760" description:"Upload file size in byte."`
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
}
