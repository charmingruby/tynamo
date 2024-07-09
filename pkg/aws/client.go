package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

const LocalStackEndpoint = "http://localhost:4566"

var (
	AwsRegion = "us-east-1"
)

func NewAWSConfig() (*aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(AwsRegion),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "tynamo", SecretAccessKey: "tynamo", SessionToken: "tynamo",
				Source: "hard-coded credentials, values are irrelevant for local DynamoDB",
			},
		}),
	)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
