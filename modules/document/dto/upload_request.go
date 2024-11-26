package dto

type UploadRequestBody struct {
	RequestSize int64 `json:"size" validate:"required" example:"10485760" description:"Upload file size in byte." min:"1" max:"10485760"`
}

type UploadDocumentRequestBody struct {
	ObjectKey string       `json:"key" validate:"required" example:"4b793c1a06ea4ea0a2b019e3c04c3f1d/c211f30fbc56484e83ca9f96afaaeb8b"`
	MetaData  MetaDataBody `json:"metadata" validate:"required"`
}

type MetaDataBody struct {
	Name      string `json:"name" validate:"required" example:"document" description:"Document name."`
	Size      int64  `json:"size" validate:"required" example:"10485760" description:"File size in byte." min:"1" max:"10485760"`
	MimeType  string `json:"mime_type" validate:"required" example:"application/pdf" description:"File MIME type."`
	Extension string `json:"extension" validate:"required" example:"pdf" description:"File extension."`
}
