package domains

import (
	"time"

	"github.com/google/uuid"

	"github.com/www-printf/wepress-core/modules/auth/domains"
)

type Document struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	ObjectKey string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	MetaData  MetaData     `gorm:"foreignKey:DocumentID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	OwnerID   uuid.UUID    `gorm:"type:uuid; not null"`
	Owner     domains.User `gorm:"foreignKey:OwnerID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type MetaData struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name       string
	Size       int64
	MimeType   string
	Extension  string
	DocumentID uuid.UUID `gorm:"type:uuid"`
}

type Tabler interface {
	TableName() string
}

func (MetaData) TableName() string {
	return "metadata"
}
