package domains

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PrintHistory struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey"`
	JobID     string    `gorm:"column:job_id"`
	PrinterID uint      `gorm:"column:printer_id"`
	ClusterID uint      `gorm:"column:cluster_id"`
}

func (p *PrintHistory) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
