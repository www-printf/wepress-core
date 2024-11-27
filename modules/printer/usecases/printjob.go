package usecases

import (
	"context"

	"github.com/www-printf/wepress-core/modules/printer/dto"
	"github.com/www-printf/wepress-core/pkg/constants"
	"github.com/www-printf/wepress-core/pkg/errors"
)

func (u *printerUsecase) SubmitPrintJob(ctx context.Context, req *dto.SubmitPrintJobRequestBody) (*dto.PrintJobResponseBody, *errors.HTTPError) {
	doc, err := u.printerRepo.GetDocument(ctx, req.DocumentID)
	if err != nil {
		return nil, constants.HTTPNotFound
	}

	// content, err := u.s3Client.DownloadObject(ctx, "", doc.ObjectKey)
	// if err != nil {
	// 	return nil, constants.HTTPInternal
	// }

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
		Content:   []byte("test"),
	}
	resp, err := u.clusterManager.SubmitPrintJob(ctx, reqJob)
	if err != nil {
		return nil, constants.HTTPInternal
	}

	return &dto.PrintJobResponseBody{
		ID:            resp.GetJobId(),
		DocumentID:    resp.GetDocumentId(),
		SubmittedAt:   resp.GetSubmittedAt().String(),
		PagesPrinted:  resp.GetPagesPrinted(),
		EstimatedTime: resp.GetEtaSeconds(),
		JobStatus:     resp.GetStatus().String(),
	}, nil
}
