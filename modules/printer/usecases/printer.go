package usecases

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/www-printf/wepress-core/modules/printer/domains"
	"github.com/www-printf/wepress-core/modules/printer/dto"
	"github.com/www-printf/wepress-core/modules/printer/repository"
	"github.com/www-printf/wepress-core/pkg/clusters"
	"github.com/www-printf/wepress-core/pkg/constants"
	"github.com/www-printf/wepress-core/pkg/errors"
	"github.com/www-printf/wepress-core/pkg/s3"
)

type PrinterUsecase interface {
	AddPrinter(ctx context.Context, req *dto.AddPrinterRequestBody) (*dto.PrinterResponseBody, *errors.HTTPError)
	GetPrinter(ctx context.Context, printerID uint) (*dto.PrinterResponseBody, *errors.HTTPError)
	ListPrinter(ctx context.Context, clusterID uint) (*dto.ListPrinterResponseBody, *errors.HTTPError)
	ViewStatus(ctx context.Context, printerID uint) (*dto.PrinterStatusResponseBody, *errors.HTTPError)
	ListCluster(ctx context.Context) (*dto.ListClusterResponseBody, *errors.HTTPError)
	SubmitPrintJob(ctx context.Context, req *dto.SubmitPrintJobRequestBody) (*dto.PrintJobResponseBody, *errors.HTTPError)
}

type printerUsecase struct {
	printerRepo    repository.PrinterRepository
	redisClient    *redis.Client
	clusterManager clusters.ClusterManager
	s3Client       s3.S3Client
}

func NewPrinterUsecase(printerRepo repository.PrinterRepository, redisClient *redis.Client, clusterManager clusters.ClusterManager, s3Client s3.S3Client) PrinterUsecase {
	return &printerUsecase{
		printerRepo:    printerRepo,
		redisClient:    redisClient,
		clusterManager: clusterManager,
		s3Client:       s3Client,
	}
}

func (u *printerUsecase) AddPrinter(ctx context.Context, req *dto.AddPrinterRequestBody) (*dto.PrinterResponseBody, *errors.HTTPError) {
	printer := &domains.Printer{
		ClusterID:    req.ClusterID,
		Manufacturer: req.Manufacturer,
		Model:        req.Model,
		SerialNumber: req.SerialNumber,
		URI:          req.URI,
	}
	printer, err := u.printerRepo.Create(ctx, printer)
	if err != nil {
		return nil, constants.HTTPInternal
	}
	return &dto.PrinterResponseBody{
		ID:           printer.ID,
		ClusterID:    printer.ClusterID,
		Manufacturer: printer.Manufacturer,
		Model:        printer.Model,
		SerialNumber: printer.SerialNumber,
		URI:          printer.URI,
		AddedAt:      printer.AddedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    printer.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (u *printerUsecase) GetPrinter(ctx context.Context, printerID uint) (*dto.PrinterResponseBody, *errors.HTTPError) {
	printer, err := u.printerRepo.GetByID(ctx, printerID)
	if err != nil {
		return nil, constants.HTTPInternal
	}
	return &dto.PrinterResponseBody{
		ID:           printer.ID,
		ClusterID:    printer.ClusterID,
		Manufacturer: printer.Manufacturer,
		Model:        printer.Model,
		SerialNumber: printer.SerialNumber,
		URI:          printer.URI,
		AddedAt:      printer.AddedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    printer.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (u *printerUsecase) ListPrinter(ctx context.Context, clusterID uint) (*dto.ListPrinterResponseBody, *errors.HTTPError) {
	printers, err := u.printerRepo.ListByClusterID(ctx, clusterID)
	if err != nil {
		return nil, constants.HTTPInternal
	}
	var printerResponses []dto.PrinterResponseBody
	for _, printer := range printers {
		printerResponses = append(printerResponses, dto.PrinterResponseBody{
			ID:           printer.ID,
			ClusterID:    printer.ClusterID,
			Manufacturer: printer.Manufacturer,
			Model:        printer.Model,
			SerialNumber: printer.SerialNumber,
			URI:          printer.URI,
			AddedAt:      printer.AddedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    printer.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	count, err := u.printerRepo.CountByClusterID(ctx, clusterID)
	if err != nil {
		return nil, constants.HTTPInternal
	}

	return &dto.ListPrinterResponseBody{
		Printers: printerResponses,
		Total:    count,
	}, nil
}

func (u *printerUsecase) ViewStatus(ctx context.Context, printerID uint) (*dto.PrinterStatusResponseBody, *errors.HTTPError) {
	
	return nil, nil
}

func (u *printerUsecase) ListCluster(ctx context.Context) (*dto.ListClusterResponseBody, *errors.HTTPError) {
	clusters, err := u.printerRepo.ListCluster(ctx)
	if err != nil {
		return nil, constants.HTTPInternal
	}

	var clusterReponses []dto.ClusterBody
	for _, cluster := range clusters {
		totalPrinter, err := u.printerRepo.CountByClusterID(ctx, cluster.ID)
		if err != nil {
			return nil, constants.HTTPInternal
		}
		clusterReponses = append(clusterReponses, dto.ClusterBody{
			ID:        cluster.ID,
			AddedAt:   cluster.AddedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: cluster.UpdatedAt.Format("2006-01-02 15:04:05"),
			Building:  cluster.Building,
			Room:      cluster.Room,
			Campus:    cluster.Campus,
			Longitude: cluster.Longitude,
			Latitude:  cluster.Latitude,
			Total:     totalPrinter,
		})
	}
	return &dto.ListClusterResponseBody{
		Clusters: clusterReponses,
		Total:    int64(len(clusters)),
	}, nil
}
