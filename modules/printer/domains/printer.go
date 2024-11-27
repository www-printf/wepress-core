package domains

import "time"

type Printer struct {
	ID           uint      `gorm:"column:id;primaryKey"`
	AddedAt      time.Time `gorm:"column:added_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
	ClusterID    uint      `gorm:"column:cluster_id"`
	Manufacturer string    `gorm:"column:manufacturer"`
	Model        string    `gorm:"column:model"`
	SerialNumber string    `gorm:"column:serial_number"`
	URI          string    `gorm:"column:uri"`
}
