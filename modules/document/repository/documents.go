package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/www-printf/wepress-core/modules/document/domains"
)

type DocumentRepository interface {
	Create(ctx context.Context, document *domains.Document) error
	GetByID(ctx context.Context, id string) (*domains.Document, error)
	GetByOwnerID(ctx context.Context, ownerID string, offset int, litmit int) ([]domains.Document, error)
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

func (r *documentRepository) GetByID(ctx context.Context, id string) (*domains.Document, error) {
	var document domains.Document
	err := r.db.WithContext(ctx).Preload("MetaData").First(&document, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &document, nil
}

func (r *documentRepository) GetByOwnerID(ctx context.Context, ownerID string, offset int, limit int) ([]domains.Document, error) {
	var documents []domains.Document
	err := r.db.WithContext(ctx).Preload("MetaData").Where("owner_id = ?", ownerID).Offset(offset).Limit(limit).Find(&documents).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return documents, nil
}
