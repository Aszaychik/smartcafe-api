package uploader

import (
	"aszaychik/smartcafe-api/config"
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AWSUploader interface {
	UploadFile(fileName, folderName string) (*manager.UploadOutput, error)
}

type AWSUploaderImpl struct {
	AWSConfig     *config.AWSConfig
	S3Client      *s3.Client
}

func NewAWSUploader(awsConfig *config.AWSConfig, s3client *s3.Client) AWSUploader {
	return &AWSUploaderImpl{
		AWSConfig:     awsConfig,
		S3Client:      s3client,
	}
}

func (awsUp *AWSUploaderImpl) UploadFile(fileName, folderName string) (*manager.UploadOutput, error) {
	uploader := manager.NewUploader(awsUp.S3Client)
	
	file, err := os.Open(fmt.Sprintf("uploads/%s", fileName))
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %w", err)
	}
	defer file.Close()

	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(awsUp.AWSConfig.BucketName),
		Key: aws.String(fmt.Sprintf("%s/%s", folderName,fileName)),
		Body: file,
		ACL: "public-read",
	})

	if err != nil {
		return nil, fmt.Errorf("failed upload : %w", err)
	}

	fmt.Println(result.Location)

	return result, nil
}