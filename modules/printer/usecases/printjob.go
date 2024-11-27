package usecases

import (
	"context"

	"github.com/www-printf/wepress-core/modules/printer/domains"
	"github.com/www-printf/wepress-core/modules/printer/dto"
	"github.com/www-printf/wepress-core/pkg/constants"
	"github.com/www-printf/wepress-core/pkg/errors"
)

func (u *printerUsecase) SubmitPrintJob(ctx context.Context, req *dto.SubmitPrintJobRequestBody) (*dto.PrintJobResponseBody, *errors.HTTPError) {
	doc, err := u.printerRepo.GetDocument(ctx, req.DocumentID)
	if err != nil {
		return nil, constants.HTTPNotFound
	}

	content, err := u.s3Client.DownloadObject(ctx, "", doc.ObjectKey)
	if err != nil {
		return nil, constants.HTTPInternal
	}

	reqJob := &dto.PrintJobTranfer{
		DocumentID: req.DocumentID,
		PrintSettings: dto.PrintSettings{
			ColorMode:   req.PrintSettings.ColorMode,
			PaperSize:   req.PrintSettings.PaperSize,
			Orientation: req.PrintSettings.Orientation,
			Copies:      req.PrintSettings.Copies,
			DoubleSided: req.PrintSettings.DoubleSided,
		},
		ClusterID: req.ClusterID,
		Name:      doc.MetaData.Name,
		Content:   content,
	}
	resp, printerID, err := u.clusterManagers[req.ClusterID].SubmitPrintJob(ctx, reqJob)
	if err != nil {
		return nil, constants.HTTPInternal
	}

	history := &domains.PrintHistory{
		JobID:     resp.GetJobId(),
		ClusterID: req.ClusterID,
		PrinterID: printerID,
	}

	if err := u.printerRepo.AddPrintHistory(ctx, history); err != nil {
		return nil, constants.HTTPInternal
	}

	return &dto.PrintJobResponseBody{
		ID:            resp.GetJobId(),
		DocumentID:    resp.GetDocumentId(),
		PagesPrinted:  resp.GetPagesPrinted(),
		EstimatedTime: resp.GetEtaSeconds(),
		JobStatus:     resp.GetStatus().String(),
		TotalPages:    resp.GetTotalPages(),
	}, nil
}

func (u *printerUsecase) ViewJobStatus(ctx context.Context, jobID string) (*dto.PrintJobResponseBody, *errors.HTTPError) {
	history, err := u.printerRepo.GetPrintHistory(ctx, jobID)
	if err != nil {
		return nil, constants.HTTPInternal
	}
	if history == nil {
		return nil, constants.HTTPNotFound
	}

	resp, err := u.clusterManagers[history.ClusterID].GetJobStatus(ctx, history.PrinterID, jobID)
	if err != nil {
		return nil, constants.HTTPInternal
	}

	return &dto.PrintJobResponseBody{
		ID:            resp.GetJobId(),
		DocumentID:    resp.GetDocumentId(),
		PagesPrinted:  resp.GetPagesPrinted(),
		EstimatedTime: resp.GetEtaSeconds(),
		JobStatus:     resp.GetStatus().String(),
		TotalPages:    resp.GetTotalPages(),
	}, nil
}

func (u *printerUsecase) CancelPrintJob(ctx context.Context, jobID string) *errors.HTTPError {
	history, err := u.printerRepo.GetPrintHistory(ctx, jobID)
	if err != nil {
		return constants.HTTPInternal
	}
	if history == nil {
		return constants.HTTPNotFound
	}

	err = u.clusterManagers[history.ClusterID].CancelPrintJob(ctx, history.PrinterID, jobID)
	if err != nil {
		return constants.HTTPInternal
	}

	return nil
}

func (u *printerUsecase) ListPrintJobs(ctx context.Context, printerID uint) (*dto.ListPrintJobResponseBody, *errors.HTTPError) {
	clusterID, err := u.printerRepo.GetClusterIDByPrinterID(ctx, printerID)
	if err != nil {
		return nil, constants.HTTPInternal
	}
	if clusterID == 0 {
		return nil, constants.HTTPNotFound
	}

	jobs, err := u.clusterManagers[clusterID].ListPrintJobs(ctx, printerID)
	if err != nil {
		return nil, constants.HTTPInternal
	}

	var resp dto.ListPrintJobResponseBody
	resp.PrintJobs = []dto.PrintJobResponseBody{}
	for _, job := range jobs.GetJobs() {
		resp.PrintJobs = append(resp.PrintJobs, dto.PrintJobResponseBody{
			ID:            job.GetJobId(),
			DocumentID:    job.GetDocumentId(),
			PagesPrinted:  job.GetPagesPrinted(),
			EstimatedTime: job.GetEtaSeconds(),
			JobStatus:     job.GetStatus().String(),
			TotalPages:    job.GetTotalPages(),
		})
	}

	return &resp, nil
}
