package s3

import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsCfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/zerolog/log"

	appCfg "github.com/www-printf/wepress-core/config"
	"github.com/www-printf/wepress-core/utils"
)

type S3Client interface {
	GeneratePostURL(ctx context.Context, bucketName, objectKey string, size int64) (*s3.PresignedPostRequest, error)
	GenerateGetURL(ctx context.Context, bucketName, objectKey string) (string, error)
	GenerateDeleteURL(ctx context.Context, bucketName, objectKey string) (string, error)
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

func (s *s3Client) GeneratePostURL(ctx context.Context, bucketName, objectKey string, size int64) (*s3.PresignedPostRequest, error) {
	if bucketName == "" {
		bucketName = s.config.BucketName
	}
	uploadSize := strconv.FormatInt(utils.Min(size, s.config.MaxSize), 10)
	return s.postObject(ctx, bucketName, objectKey, uploadSize)
}

func (s *s3Client) GenerateGetURL(ctx context.Context, bucketName, objectKey string) (string, error) {
	if bucketName == "" {
		bucketName = s.config.BucketName
	}
	return s.getObject(ctx, bucketName, objectKey)
}

func (s *s3Client) GenerateDeleteURL(ctx context.Context, bucketName, objectKey string) (string, error) {
	if bucketName == "" {
		bucketName = s.config.BucketName
	}
	return s.deleteObject(ctx, bucketName, objectKey)
}

func (s *s3Client) postObject(ctx context.Context, bucketName, objectKey string, size string) (*s3.PresignedPostRequest, error) {
	request, err := s.presigner.PresignPostObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}, func(opts *s3.PresignPostOptions) {
		opts.Expires = time.Duration(s.config.PresignedExpire * int(time.Second))
		opts.Conditions = []interface{}{
			[]string{"content-length-range", "0", size},
		}
	})
	if err != nil {
		return nil, err
	}
	return request, err
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
