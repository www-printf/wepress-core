package domains

import "time"

type Cluster struct {
	ID        uint      `gorm:"column:id;primaryKey"`
	Status    string    `gorm:"column:status"`
	AddedAt   time.Time `gorm:"column:added_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Building  string    `gorm:"column:building"`
	Room      string    `gorm:"column:room"`
	Campus    string    `gorm:"column:campus"`
	Longitude float64   `gorm:"column:longitude"`
	Latitude  float64   `gorm:"column:latitude"`

	Printers []Printer `gorm:"foreignKey:ClusterID"`
}

type Printer struct {
	ID           uint      `gorm:"column:id;primaryKey"`
	Status       string    `gorm:"column:status"`
	AddedAt      time.Time `gorm:"column:added_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
	ClusterID    uint      `gorm:"column:cluster_id"`
	Manufacturer string    `gorm:"column:manufacturer"`
	Model        string    `gorm:"column:model"`
	SerialNumber string    `gorm:"column:serial_number"`
	IPAddress    string    `gorm:"column:ip_address"`
	MACAddress   string    `gorm:"column:mac_address"`
}

const (
	PrinterStatusActive   string = "active"
	PrinterStatusInactive string = "inactive"
)
