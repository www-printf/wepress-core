package dto

type PrinterResponseBody struct {
	ID           uint   `json:"id"`
	ClusterID    uint   `json:"cluster_id"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	SerialNumber string `json:"serial_number"`
	IPAddress    string `json:"ip_address"`
	MACAddress   string `json:"mac_address"`
	Status       string `json:"status"`
	AddedAt      string `json:"added_at"`
	UpdatedAt    string `json:"updated_at"`
}

type ListPrinterResponseBody struct {
	Printers []PrinterResponseBody `json:"printers"`
	Total    int64                 `json:"total_printers"`
	Active   int64                 `json:"active_printers"`
	Inactive int64                 `json:"inactive_printers"`
}

type PrinterStatusResponseBody struct {
	ID     uint   `json:"id"`
	Status string `json:"status"`
}
