package usecases

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/www-printf/wepress-core/modules/document/domains"
	"github.com/www-printf/wepress-core/modules/document/dto"
	"github.com/www-printf/wepress-core/modules/document/repository"
	"github.com/www-printf/wepress-core/pkg/constants"
	"github.com/www-printf/wepress-core/pkg/errors"
	"github.com/www-printf/wepress-core/pkg/s3"
)

type DocumentUsecase interface {
	RequestUpload(ctx context.Context, req *dto.UploadRequestBody, uid string) (*dto.PresignedURLResponseBody, *errors.HTTPError)
	SaveDocument(ctx context.Context, req *dto.UploadDocumentRequestBody, uid string) *errors.HTTPError
}

type documentUsecase struct {
	documentRepo repository.DocumentRepository
	s3Client     s3.S3Client
}

func NewDocumentUsecase(
	documentRepo repository.DocumentRepository,
	s3Client s3.S3Client,
) DocumentUsecase {
	return &documentUsecase{
		documentRepo: documentRepo,
		s3Client:     s3Client,
	}
}

func (u *documentUsecase) RequestUpload(ctx context.Context, req *dto.UploadRequestBody, uid string) (*dto.PresignedURLResponseBody, *errors.HTTPError) {
	objectKey := u.generateObjectKey(uid)
	presigned, err := u.s3Client.GeneratePostURL(ctx, "", objectKey, req.RequestSize)
	if err != nil {
		return nil, constants.HTTPInternal
	}
	return &dto.PresignedURLResponseBody{
		URL:       presigned.URL,
		ObjectKey: objectKey,
		Fields:    presigned.Values,
	}, nil
}

func (u *documentUsecase) SaveDocument(ctx context.Context, req *dto.UploadDocumentRequestBody, uid string) *errors.HTTPError {
	userID, err := uuid.Parse(uid)
	if err != nil {
		return constants.HTTPUnauthorized
	}
	docID := uuid.New()
	doc := &domains.Document{
		ID:        docID,
		OwnerID:   userID,
		ObjectKey: req.ObjectKey,
		MetaData: domains.MetaData{
			ID:         uuid.New(),
			Name:       req.MetaData.Name,
			Size:       req.MetaData.Size,
			MimeType:   req.MetaData.MimeType,
			Extension:  req.MetaData.Extension,
			DocumentID: docID,
		},
	}
	err = u.documentRepo.Create(ctx, doc)
	if err != nil {
		return constants.HTTPInternal
	}
	return nil
}

func (u *documentUsecase) generateObjectKey(userID string) string {
	name := uuid.New().String()
	return fmt.Sprintf("%s/%s", userID, name)
}
