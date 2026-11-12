package s3

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"profile-template/internal/assets"
	"profile-template/internal/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3Storage struct {
	client *s3.Client
	bucket string
}

func NewS3Storage(
	ctx context.Context,
	conf config.StorageConfig,
) (assets.RepositoryContract, error) {

	cfg, err := awsconfig.LoadDefaultConfig(
		ctx,
		awsconfig.WithRegion(conf.Region),
		awsconfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				conf.AccessKey,
				conf.SecretKey,
				"",
			),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("load s3 config: %w", err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(conf.Endpoint)
		o.UsePathStyle = true
	})

	// Validate bucket availability.
	_, err = client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(conf.Bucket),
	})
	if err != nil {
		return nil, fmt.Errorf(
			"%w: bucket unavailable: %v",
			assets.ErrStorageUnavailable,
			err,
		)
	}

	return &S3Storage{
		client: client,
		bucket: conf.Bucket,
	}, nil
}

func (s *S3Storage) Write(
	ctx context.Context,
	reader io.Reader,
	destination string,
) (assets.StoreResult, error) {

	key, err := resolveObjectKey(destination)
	if err != nil {
		return assets.StoreResult{}, assets.ErrAssetInvalidPath
	}

	// Read all data to determine exact content length.
	data, err := io.ReadAll(reader)
	if err != nil || len(data) == 0 {
		return assets.StoreResult{}, assets.ErrAssetFailedCreate
	}

	_, err = s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:        aws.String(s.bucket),
		Key:           aws.String(key),
		Body:          bytes.NewReader(data),
		ContentLength: aws.Int64(int64(len(data))),
	})
	if err != nil {
		return assets.StoreResult{}, fmt.Errorf(
			"put object: %w", err,
		)
	}

	return assets.StoreResult{
		FileKey: key,
		Size:    int64(len(data)),
	}, nil
}

func (s *S3Storage) Read(
	ctx context.Context,
	destination string,
) (io.ReadCloser, error) {

	key, err := resolveObjectKey(destination)
	if err != nil {
		return nil, assets.ErrAssetInvalidPath
	}

	result, err := s.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		var notFound *types.NoSuchKey
		if errors.As(err, &notFound) {
			return nil, assets.ErrAssetNotFound
		}

		return nil, fmt.Errorf("get object: %w", err)
	}

	return result.Body, nil
}

func (s *S3Storage) Delete(
	ctx context.Context,
	destination string,
) error {

	key, err := resolveObjectKey(destination)
	if err != nil {
		return assets.ErrAssetInvalidPath
	}

	_, err = s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("delete object: %w", err)
	}

	return nil
}
