package usecases

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"

	"github.com/www-printf/wepress-core/modules/printer/domains"
	"github.com/www-printf/wepress-core/modules/printer/dto"
	"github.com/www-printf/wepress-core/modules/printer/repository"
	"github.com/www-printf/wepress-core/pkg/constants"
	"github.com/www-printf/wepress-core/pkg/errors"
)

type PrinterUsecase interface {
	AddPrinter(ctx context.Context, req *dto.AddPrinterRequestBody) (*dto.PrinterResponseBody, *errors.HTTPError)
	GetPrinter(ctx context.Context, printerID uint) (*dto.PrinterResponseBody, *errors.HTTPError)
	ListPrinter(ctx context.Context, clusterID uint) (*dto.ListPrinterResponseBody, *errors.HTTPError)
	ViewStatus(ctx context.Context, printerID uint) (*dto.PrinterStatusResponseBody, *errors.HTTPError)
}

type printerUsecase struct {
	printerRepo repository.PrinterRepository
	redisClient *redis.Client
}

func NewPrinterUsecase(printerRepo repository.PrinterRepository, redisClient *redis.Client) PrinterUsecase {
	return &printerUsecase{
		printerRepo: printerRepo,
		redisClient: redisClient,
	}
}

func (u *printerUsecase) AddPrinter(ctx context.Context, req *dto.AddPrinterRequestBody) (*dto.PrinterResponseBody, *errors.HTTPError) {
	printer := &domains.Printer{
		ClusterID:    req.ClusterID,
		Manufacturer: req.Manufacturer,
		Model:        req.Model,
		SerialNumber: req.SerialNumber,
		IPAddress:    req.IPAddress,
		MACAddress:   req.MACAddress,
		Status:       req.Status,
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
		IPAddress:    printer.IPAddress,
		MACAddress:   printer.MACAddress,
		Status:       printer.Status,
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
		IPAddress:    printer.IPAddress,
		MACAddress:   printer.MACAddress,
		Status:       printer.Status,
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
			IPAddress:    printer.IPAddress,
			MACAddress:   printer.MACAddress,
			Status:       printer.Status,
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
	status, err := u.redisClient.Get(ctx, strconv.FormatUint(uint64(printerID), 10)).Result()
	if err != nil && err != redis.Nil {
		log.Error().Err(err).Msg("1")
		return nil, constants.HTTPInternal
	}
	if err == redis.Nil {
		printer, err := u.printerRepo.GetByID(ctx, printerID)
		if err != nil {
			log.Error().Err(err).Msg("2")
			return nil, constants.HTTPInternal
		}
		status = printer.Status
		u.redisClient.Set(ctx, strconv.FormatUint(uint64(printerID), 10), status, 1*time.Hour)
	}
	return &dto.PrinterStatusResponseBody{
		ID:     printerID,
		Status: status,
	}, nil
}
