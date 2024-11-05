package domains

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	FullName  string    `gorm:"not null;column:fullname"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	PrivKey   string `gorm:"not null;column:privkey"`
	PubKey    string `gorm:"not null;column:pubkey"`
}
