package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewDynamoConnection(cfg *aws.Config) *dynamodb.Client {
	resolver := dynamodb.NewDefaultEndpointResolverV2()
	resolver.ResolveEndpoint(context.TODO(), dynamodb.EndpointParameters{
		Region:   &cfg.Region,
		Endpoint: aws.String(LocalStackEndpoint),
	})

	svc := dynamodb.NewFromConfig(*cfg, func(o *dynamodb.Options) {
		o.EndpointResolverV2 = resolver
	})

	return svc
}
