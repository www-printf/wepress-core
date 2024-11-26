package domains

import "time"

type Cluster struct {
	ID        uint      `gorm:"column:id;primaryKey"`
	AddedAt   time.Time `gorm:"column:added_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Building  string    `gorm:"column:building"`
	Room      string    `gorm:"column:room"`
	Campus    string    `gorm:"column:campus"`
	Longitude float64   `gorm:"column:longitude"`
	Latitude  float64   `gorm:"column:latitude"`

	Printers []Printer `gorm:"foreignKey:ClusterID"`
}
