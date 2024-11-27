package domains

import "github.com/google/uuid"

type PrintHistory struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey"`
	JobID     string    `gorm:"column:job_id"`
	PrinterID uint      `gorm:"column:printer_id"`
	ClusterID uint      `gorm:"column:cluster_id"`
}
