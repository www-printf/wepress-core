package domains

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	FullName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
