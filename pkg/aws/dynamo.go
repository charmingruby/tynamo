package aws

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewDynamoConnection(cfg *aws.Config) (*dynamodb.Client, error) {
	svc := dynamodb.NewFromConfig(*cfg, func(o *dynamodb.Options) {
		o.BaseEndpoint = aws.String(LocalStackEndpoint)
	})

	response, err := svc.ListTables(context.TODO(), &dynamodb.ListTablesInput{})
	if err != nil {
		return nil, err
	}

	if len(response.TableNames) == 0 {
		slog.Info("No table found")
	} else {
		for _, tableName := range response.TableNames {
			slog.Info(fmt.Sprintf("Table found: %s", tableName))
		}
	}

	return svc, nil
}
