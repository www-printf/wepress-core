package dto

type ClusterBody struct {
	ID        uint    `json:"id"`
	Status    string  `json:"status"`
	AddedAt   string  `json:"added_at"`
	UpdatedAt string  `json:"updated_at"`
	Building  string  `json:"building"`
	Room      string  `json:"room"`
	Campus    string  `json:"campus"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Total     int64   `json:"total_printers"`
}

type ListClusterResponseBody struct {
	Clusters []ClusterBody `json:"clusters"`
	Total    int64         `json:"total_clusters"`
}
