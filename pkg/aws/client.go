package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

const LocalStackEndpoint = "http://localhost:4566"

var (
	awsRegion    = "us-east-1"
	awsAccessKey = "tynamo"
	awsSecretKey = "tynamo"
)

func NewAWSConfig() (*aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsAccessKey, awsSecretKey, "")),
	)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
