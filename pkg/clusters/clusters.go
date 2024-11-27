package clusters

import (
	"context"
	"errors"
	"math"

	"github.com/www-printf/wepress-core/modules/printer/dto"
	"github.com/www-printf/wepress-core/modules/printer/proto"
	"github.com/www-printf/wepress-core/utils"
	"google.golang.org/grpc"
)

type ClusterManager interface {
	AddPrinterClient(id uint, client proto.VirtualPrinterClient, conn *grpc.ClientConn)
	SubmitPrintJob(ctx context.Context, reqJob *dto.PrintJobTranfer) (*proto.PrintJob, uint, error)
	GetJobStatus(ctx context.Context, printerID uint, jobID string) (*proto.PrintJob, error)
	CancelPrintJob(ctx context.Context, printerID uint, jobID string) error
	Close()
}

type clusterManager struct {
	printers map[uint]proto.VirtualPrinterClient
	conns    map[uint]*grpc.ClientConn
}

func NewClusterManager() ClusterManager {
	return &clusterManager{
		printers: make(map[uint]proto.VirtualPrinterClient),
		conns:    make(map[uint]*grpc.ClientConn),
	}
}

func (m *clusterManager) AddPrinterClient(id uint, client proto.VirtualPrinterClient, conn *grpc.ClientConn) {
	m.printers[id] = client
	m.conns[id] = conn
}

func (m *clusterManager) Close() {
	for _, conn := range m.conns {
		conn.Close()
	}
}

func (m *clusterManager) SubmitPrintJob(ctx context.Context, reqJob *dto.PrintJobTranfer) (*proto.PrintJob, uint, error) {
	bestPrinter, err := m.findBestPrinter(ctx)
	if err != nil {
		return nil, 0, err
	}
	printer := m.printers[bestPrinter]
	resp, err := printer.SubmitPrintJob(ctx, &proto.PrintDocument{
		DocumentId: reqJob.DocumentID,
		Name:       reqJob.Name,
		Content:    reqJob.Content,
		Settings: &proto.PrintSettings{
			ColorMode:   utils.MapColorMode(reqJob.PrintSettings.ColorMode),
			PaperSize:   utils.MapPaperSize(reqJob.PrintSettings.PaperSize),
			Orientation: utils.MapOrientation(reqJob.PrintSettings.Orientation),
			Copies:      reqJob.PrintSettings.Copies,
			DoubleSided: reqJob.PrintSettings.DoubleSided,
		}})
	if err != nil {
		return nil, 0, err
	}

	return resp, bestPrinter, nil
}

func (m *clusterManager) GetJobStatus(ctx context.Context, printerID uint, jobID string) (*proto.PrintJob, error) {
	printer := m.printers[printerID]
	resp, err := printer.GetJobStatus(ctx, &proto.GetJobStatusRequest{JobId: jobID})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *clusterManager) CancelPrintJob(ctx context.Context, printerID uint, jobID string) error {
	printer := m.printers[printerID]
	_, err := printer.CancelPrintJob(ctx, &proto.CancelJobRequest{JobId: jobID})
	if err != nil {
		return err
	}
	return nil
}

func (m *clusterManager) findBestPrinter(ctx context.Context) (uint, error) {
	var choosenPrinter uint
	printerFound := false
	minEstimatedTime := int32(math.MaxInt32)
	for id, printer := range m.printers {
		queuedJobs, err := printer.ListPrintJobs(ctx, &proto.Empty{})
		if err != nil {
			continue
		}
		jobs := queuedJobs.GetJobs()
		totalEstimatedTime := int32(0)
		for _, job := range jobs {
			totalEstimatedTime += job.EtaSeconds
		}
		if totalEstimatedTime < minEstimatedTime && len(jobs) < 10 {
			choosenPrinter = id
			minEstimatedTime = totalEstimatedTime
			printerFound = true
		}
	}
	if !printerFound {
		return 0, errors.New("no available printer")
	}
	return choosenPrinter, nil
}
