package dynamo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/charmingruby/tynamo/internal/model"
	"github.com/charmingruby/tynamo/internal/repository/dynamo/mapper"
)

func NewDynamoExampleRepository(db *dynamodb.Client) *DynamoExampleRepository {
	return &DynamoExampleRepository{
		db: db,
	}
}

type DynamoExampleRepository struct {
	db *dynamodb.Client
}

func (r *DynamoExampleRepository) Store(example *model.Example) error {
	data, err := mapper.ExampleModelToDynamo(example)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("examples"),
		Item:      data,
	}

	_, err = r.db.PutItem(context.TODO(), input)
	if err != nil {
		return err
	}

	return nil
}

func (r *DynamoExampleRepository) List() ([]model.Example, error) {

	input := &dynamodb.ScanInput{
		TableName: aws.String("examples"),
	}

	result, err := r.db.Scan(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	var examples []model.Example
	for _, item := range result.Items {
		examples = append(examples, mapper.ExampleDynamoToModel(item))
	}

	return examples, nil
}
