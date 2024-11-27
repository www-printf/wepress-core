package domains

import "time"

type PrintJob struct {
	ID          string    `gorm:"primaryKey"`
	DocumentID  string    `gorm:"not null"`
	SubmittedAt time.Time `gorm:"not null"`
	TotalPages  int       `gorm:"not null"`
}
