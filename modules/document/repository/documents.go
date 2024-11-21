package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/www-printf/wepress-core/modules/document/domains"
)

type DocumentRepository interface {
	Create(ctx context.Context, document *domains.Document) error
}

type documentRepository struct {
	db *gorm.DB
}

func NewDocumentRepository(db *gorm.DB) DocumentRepository {
	return &documentRepository{db: db}
}

func (r *documentRepository) Create(ctx context.Context, document *domains.Document) error {
	err := r.db.WithContext(ctx).Create(document).Error
	if err != nil {
		return err
	}
	return nil
}
