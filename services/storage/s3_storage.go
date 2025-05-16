package storage

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Storage struct {
	client *s3.Client
}

func NewS3Storage() *S3Storage {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(os.Getenv("AWS_REGION")),
	)
	if err != nil {
		log.Fatal(err)
	}
	client := s3.NewFromConfig(cfg)
	return &S3Storage{client: client}
}

func (s *S3Storage) Upload(fileName string, file io.Reader, contentType string) (string, error) {
	_, err := s.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String("aasksar"),
		Key:         aws.String(fileName),
		Body:        file,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", err 
	}

	return fileName, nil
}

func (s *S3Storage) Download(key string) (io.Reader, error) {
	o, err := s.client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String("aasksar"),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	return o.Body, nil
}

func (s *S3Storage) ListAll() (map[string]string, error) {
	output, err := s.client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("aasksar"),
	})
	if err != nil {
		return nil, err
	}
	objectsMap := make(map[string]string)
	for _, obj := range output.Contents {
		objectsMap[*obj.Key] = *obj.ETag
	}
	return objectsMap, nil
}
