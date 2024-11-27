package clusters

import (
	"context"
	"errors"
	"math"
	"time"

	"github.com/www-printf/wepress-core/modules/printer/dto"
	"github.com/www-printf/wepress-core/modules/printer/proto"
	"github.com/www-printf/wepress-core/modules/printer/repository"
	"github.com/www-printf/wepress-core/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClusterManager interface {
	SubmitPrintJob(ctx context.Context, reqJob *dto.PrintJobTranfer) (*proto.PrintJob, error)
	Close()
}

type clusterManager struct {
	clusters map[uint]*Cluster
	repo     repository.PrinterRepository
}

type Cluster struct {
	printers map[uint]proto.VirtualPrinterClient
	conns    map[uint]*grpc.ClientConn
}

func NewClusterManager(repo repository.PrinterRepository) ClusterManager {
	manager := &clusterManager{
		clusters: make(map[uint]*Cluster),
		repo:     repo,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	clustersFromDB, err := repo.ListCluster(ctx)
	if err != nil {
		return nil
	}

	for _, cluster := range clustersFromDB {
		printers := make(map[uint]proto.VirtualPrinterClient)
		conns := make(map[uint]*grpc.ClientConn)
		for _, printer := range cluster.Printers {
			conn, client, err := newPrinterClient(printer.URI)
			if err != nil {
				continue
			}
			printers[printer.ID] = client
			conns[printer.ID] = conn
		}
		manager.clusters[cluster.ID] = &Cluster{
			printers: printers,
			conns:    conns,
		}
	}

	return manager
}

func newPrinterClient(uri string) (*grpc.ClientConn, proto.VirtualPrinterClient, error) {
	conn, err := grpc.NewClient(uri, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	client := proto.NewVirtualPrinterClient(conn)
	return conn, client, nil
}

func (m *clusterManager) Close() {
	for _, cluster := range m.clusters {
		for _, conn := range cluster.conns {
			conn.Close()
		}
	}
}

func (m *clusterManager) SubmitPrintJob(ctx context.Context, reqJob *dto.PrintJobTranfer) (*proto.PrintJob, error) {
	cluster, ok := m.clusters[reqJob.ClusterID]
	if !ok {
		return nil, errors.New("cluster not found")
	}
	var choosenPrinter uint
	printerFound := false
	minEstimatedTime := int32(math.MaxInt32)
	for id, printer := range cluster.printers {
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
		return nil, errors.New("no available printer")
	}
	printer := cluster.printers[choosenPrinter]
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
		return nil, err
	}
	return resp, nil
}
