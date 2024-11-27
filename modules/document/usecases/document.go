package usecases

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"github.com/www-printf/wepress-core/modules/document/domains"
	"github.com/www-printf/wepress-core/modules/document/dto"
	"github.com/www-printf/wepress-core/modules/document/repository"
	"github.com/www-printf/wepress-core/pkg/constants"
	"github.com/www-printf/wepress-core/pkg/errors"
	"github.com/www-printf/wepress-core/pkg/s3"
)

type DocumentUsecase interface {
	RequestUpload(ctx context.Context, req *dto.UploadRequestBody, uid string) (*dto.RequestUploadResponseBody, *errors.HTTPError)
	SaveDocument(ctx context.Context, req *dto.UploadDocumentRequestBody, uid string) (*dto.UploadResponseBody, *errors.HTTPError)
	DownloadDocument(ctx context.Context, docID string, uid string) (*dto.DownloadDocumentResponseBody, *errors.HTTPError)
	DownloadDocumentList(ctx context.Context, uid string, page int, perPage int) (*dto.DownloadDocumentsResponseBody, *errors.HTTPError)
}

type documentUsecase struct {
	s3Client     s3.S3Client
	redisClient  *redis.Client
	documentRepo repository.DocumentRepository
}

func NewDocumentUsecase(
	documentRepo repository.DocumentRepository,
	s3Client s3.S3Client,
	redisClient *redis.Client,
) DocumentUsecase {
	return &documentUsecase{
		documentRepo: documentRepo,
		s3Client:     s3Client,
		redisClient:  redisClient,
	}
}

func (u *documentUsecase) RequestUpload(ctx context.Context, req *dto.UploadRequestBody, uid string) (*dto.RequestUploadResponseBody, *errors.HTTPError) {
	objectKey := u.generateObjectKey(uid)
	presigned, err := u.s3Client.GeneratePostURL(ctx, u.s3Client.GetConfig().BucketName, objectKey, req.RequestSize)
	if err != nil {
		return nil, constants.HTTPInternal
	}
	err = u.setUploadStatus(ctx, objectKey)
	if err != nil {
		return nil, constants.HTTPInternal
	}
	return &dto.RequestUploadResponseBody{
		URL:       presigned.URL,
		ObjectKey: objectKey,
		Fields:    presigned.Values,
	}, nil
}

func (u *documentUsecase) SaveDocument(ctx context.Context, req *dto.UploadDocumentRequestBody, uid string) (*dto.UploadResponseBody, *errors.HTTPError) {
	userID, err := uuid.Parse(uid)
	if err != nil {
		return nil, constants.HTTPUnauthorized
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
	isOwner := u.checkOwner(uid, req.ObjectKey)
	if !isOwner {
		return nil, constants.HTTPForbidden
	}
	uploadStatus, err := u.checkUploadStatus(ctx, req.ObjectKey)
	if err != nil {
		return nil, constants.HTTPInternal
	}
	if !uploadStatus {
		return nil, constants.HTTPBadRequest
	}
	err = u.documentRepo.Create(ctx, doc)
	if err != nil {
		return nil, constants.HTTPInternal
	}
	return &dto.UploadResponseBody{
		ID:       doc.ID.String(),
		MetaData: req.MetaData,
	}, nil
}

func (u *documentUsecase) DownloadDocument(ctx context.Context, docID string, uid string) (*dto.DownloadDocumentResponseBody, *errors.HTTPError) {
	doc, err := u.documentRepo.GetByID(ctx, docID)
	if err != nil {
		return nil, constants.HTTPInternal
	}
	if doc == nil {
		return nil, constants.HTTPNotFound
	}
	isOwner := doc.OwnerID.String() == uid
	if !isOwner {
		return nil, constants.HTTPForbidden
	}
	url, err := u.s3Client.GenerateGetURL(ctx, u.s3Client.GetConfig().BucketName, doc.ObjectKey)
	if err != nil {
		return nil, constants.HTTPInternal
	}
	return &dto.DownloadDocumentResponseBody{
		URL: url,
		ID:  docID,
		MetaData: dto.MetaDataBody{
			Name:      doc.MetaData.Name,
			Size:      doc.MetaData.Size,
			MimeType:  doc.MetaData.MimeType,
			Extension: doc.MetaData.Extension,
		},
		CreatedAt: doc.CreatedAt.Format(time.RFC3339),
		UpdatedAt: doc.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (u *documentUsecase) DownloadDocumentList(ctx context.Context, uid string, page int, perPage int) (*dto.DownloadDocumentsResponseBody, *errors.HTTPError) {
	if perPage == 0 {
		perPage = constants.DefaultPerPage
	}
	if perPage > constants.MaximumPerPage {
		perPage = constants.MaximumPerPage
	}
	offset := (page - 1) * perPage
	docs, err := u.documentRepo.GetByOwnerID(ctx, uid, offset, perPage)
	if err != nil {
		return nil, constants.HTTPInternal
	}
	var downloadDocs []dto.DownloadDocumentResponseBody
	for _, doc := range docs {
		url, err := u.s3Client.GenerateGetURL(ctx, u.s3Client.GetConfig().BucketName, doc.ObjectKey)
		if err != nil {
			return nil, constants.HTTPInternal
		}
		downloadDocs = append(downloadDocs, dto.DownloadDocumentResponseBody{
			URL: url,
			ID:  doc.ID.String(),
			MetaData: dto.MetaDataBody{
				Name:      doc.MetaData.Name,
				Size:      doc.MetaData.Size,
				MimeType:  doc.MetaData.MimeType,
				Extension: doc.MetaData.Extension,
			},
			CreatedAt: doc.CreatedAt.Format(time.RFC3339),
			UpdatedAt: doc.UpdatedAt.Format(time.RFC3339),
		})
	}
	return &dto.DownloadDocumentsResponseBody{
		Documents: downloadDocs,
	}, nil
}

func (u *documentUsecase) generateObjectKey(userID string) string {
	name := strings.ReplaceAll(uuid.New().String(), "-", "")
	userID = strings.ReplaceAll(userID, "-", "")
	return fmt.Sprintf("%s/%s", userID, name)
}

func (u *documentUsecase) checkUploadStatus(ctx context.Context, objectKey string) (bool, error) {
	_, err := u.redisClient.Get(ctx, objectKey).Result()
	switch {
	case err == redis.Nil:
		return false, nil
	case err != nil:
		return false, err
	default:
		return true, nil
	}
}

func (u *documentUsecase) setUploadStatus(ctx context.Context, objectKey string) error {
	expriation := time.Duration(u.s3Client.GetConfig().PresignedExpire) * time.Second
	return u.redisClient.Set(ctx, objectKey, "1", expriation).Err()
}

func (u *documentUsecase) checkOwner(userID string, objectKey string) bool {
	prefixUID := strings.Split(objectKey, "/")[0]
	return strings.ReplaceAll(userID, "-", "") == prefixUID
}
