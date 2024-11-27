package dto

type AddPrinterRequestBody struct {
	ClusterID    uint   `json:"cluster_id"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	SerialNumber string `json:"serial_number"`
	URI          string `json:"uri"`
	Status       string `json:"status"`
}