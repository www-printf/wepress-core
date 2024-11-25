package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/www-printf/wepress-core/modules/printer/domains"
)

type PrinterRepository interface {
	Create(ctx context.Context, printer *domains.Printer) (*domains.Printer, error)
	GetByID(ctx context.Context, id uint) (*domains.Printer, error)
	ListByClusterID(ctx context.Context, clusterID uint) ([]domains.Printer, error)
	CountByClusterID(ctx context.Context, clusterID uint) (int64, error)
	CountActiveByClusterID(ctx context.Context, clusterID uint) (int64, error)
	ListCluster(ctx context.Context) ([]domains.Cluster, error)
}

type printerRepository struct {
	db *gorm.DB
}

func NewPrinterRepository(db *gorm.DB) PrinterRepository {
	return &printerRepository{db: db}
}

func (r *printerRepository) Create(ctx context.Context, printer *domains.Printer) (*domains.Printer, error) {
	err := r.db.WithContext(ctx).Create(printer).Error
	if err != nil {
		return nil, err
	}
	return printer, nil
}

func (r *printerRepository) GetByID(ctx context.Context, id uint) (*domains.Printer, error) {
	var printer domains.Printer
	err := r.db.WithContext(ctx).First(&printer, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &printer, nil
}

func (r *printerRepository) ListByClusterID(ctx context.Context, clusterID uint) ([]domains.Printer, error) {
	var printers []domains.Printer
	err := r.db.WithContext(ctx).Where("cluster_id = ?", clusterID).Find(&printers).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return printers, nil
}

func (r *printerRepository) CountByClusterID(ctx context.Context, clusterID uint) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&domains.Printer{}).Where("cluster_id = ?", clusterID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *printerRepository) CountActiveByClusterID(ctx context.Context, clusterID uint) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&domains.Printer{}).Where("cluster_id = ? AND status = ?", clusterID, domains.PrinterStatusActive).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *printerRepository) ListCluster(ctx context.Context) ([]domains.Cluster, error) {
	var clusters []domains.Cluster
	err := r.db.WithContext(ctx).Find(&clusters).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return clusters, nil
}
