package dto

type PrinterResponseBody struct {
	ID           uint   `json:"id"`
	ClusterID    uint   `json:"cluster_id"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	SerialNumber string `json:"serial_number"`
	URI          string `json:"uri"`
	AddedAt      string `json:"added_at"`
	UpdatedAt    string `json:"updated_at"`
}

type ListPrinterResponseBody struct {
	Printers []PrinterResponseBody `json:"printers"`
	Total    int64                 `json:"total_printers"`
}

type PrinterStatusResponseBody struct {
	Status PrinterStatus         `json:"status"`
	Job    *PrintJobResponseBody `json:"running_job"`
}

type PrinterStatus string

const (
	PrinterStatusIdle        PrinterStatus = "idle"
	PrinterStatusPrinting    PrinterStatus = "printing"
	PrinterStatusError       PrinterStatus = "error"
	PrinterStatusUnspecified PrinterStatus = "unspecified"
)
