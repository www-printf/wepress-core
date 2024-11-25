package dto

type AddPrinterRequestBody struct {
	ClusterID    uint   `json:"cluster_id"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	SerialNumber string `json:"serial_number"`
	IPAddress    string `json:"ip_address"`
	MACAddress   string `json:"mac_address"`
	Status       string `json:"status"`
}
