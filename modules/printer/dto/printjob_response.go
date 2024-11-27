package dto

type PrintJobResponseBody struct {
	ID            string `json:"id"`
	DocumentID    string `json:"document_id"`
	SubmittedAt   string `json:"submitted_at"`
	PagesPrinted  int32  `json:"pages_printed"`
	TotalPages    int32  `json:"total_pages"`
	EstimatedTime int32  `json:"estimated_time"`
	JobStatus     string `json:"job_status"`
}
