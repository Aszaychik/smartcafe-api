package aws

import (
	"context"
	"fmt"

	cfg "aszaychik/smartcafe-api/config"

	"github.com/aws/aws-sdk-go-v2/config"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewS3Client(awsConfig *cfg.AWSConfig) (*s3.Client, error) {
	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(awsConfig.Region))
	if err != nil {
		return nil, fmt.Errorf("Error loading AWS SDK configuration: %w", err)
	}
	client := s3.NewFromConfig(config)

	return client, nil
}