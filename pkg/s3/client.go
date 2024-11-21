package s3

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsCfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/zerolog/log"

	appCfg "github.com/www-printf/wepress-core/config"
	"github.com/www-printf/wepress-core/pkg/constants"
)

type S3Client interface {
	GeneratePresignedURL(ctx context.Context, bucketName, objectKey, action string) (string, error)
}

type s3Client struct {
	config    *appCfg.S3Config
	presigner *s3.PresignClient
}

func NewS3Client(appConf *appCfg.AppConfig) S3Client {
	creds := credentials.NewStaticCredentialsProvider(
		appConf.S3Config.AccessKey,
		appConf.S3Config.SecretKey,
		"",
	)

	cfg, err := awsCfg.LoadDefaultConfig(
		context.TODO(),
		awsCfg.WithRegion(appConf.S3Config.Region),
		awsCfg.WithCredentialsProvider(creds),
	)

	if err != nil {
		log.Error().Err(err).Msg("failed to load aws config")
		return nil
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(appConf.S3Config.EndPoint)
	})

	return &s3Client{
		config:    &appConf.S3Config,
		presigner: s3.NewPresignClient(client),
	}
}

func (s *s3Client) GeneratePresignedURL(ctx context.Context, bucketName, objectKey, action string) (string, error) {
	if bucketName == "" {
		bucketName = s.config.BucketName
	}
	switch action {
	case constants.S3ActionDownload:
		url, err := s.getObject(ctx, bucketName, objectKey)
		if err != nil {
			return "", err
		}
		return url, nil
	case constants.S3ActionUpload:
		url, err := s.putObject(ctx, bucketName, objectKey)
		if err != nil {
			return "", err
		}
		return url, nil
	case constants.S3ActionDelete:
		url, err := s.deleteObject(ctx, bucketName, objectKey)
		if err != nil {
			return "", err
		}
		return url, nil
	default:
		return "", errors.New("invalid action")
	}

}

func (s *s3Client) putObject(ctx context.Context, bucketName, objectKey string) (string, error) {
	request, err := s.presigner.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(s.config.PresignedExpire * int(time.Second))
	})
	if err != nil {
		return "", err
	}
	return request.URL, err
}

func (s *s3Client) getObject(ctx context.Context, bucketName, objectKey string) (string, error) {
	request, err := s.presigner.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(s.config.PresignedExpire * int(time.Second))
	})
	if err != nil {
		return "", err
	}
	return request.URL, err
}

func (s *s3Client) deleteObject(ctx context.Context, bucketName, objectKey string) (string, error) {
	request, err := s.presigner.PresignDeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return "", err
	}
	return request.URL, err
}
