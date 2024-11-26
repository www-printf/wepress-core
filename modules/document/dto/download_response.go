package dto

type DownloadDocumentResponseBody struct {
	ID       string       `json:"id" validate:"required" example:"d2728e88-aef1-4822-976a-63bdca2e89f9"`
	URL      string       `json:"url" validate:"required" example:"https://bucket.s3-endpoint/object-key" description:"Download URL."`
	MetaData MetaDataBody `json:"metadata" validate:"required"`
}

type DownloadDocumentsResponseBody struct {
	Documents []DownloadDocumentResponseBody `json:"documents" validate:"required"`
}
