package mapper

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/charmingruby/tynamo/internal/model"
)

func ExampleModelToDynamo(e *model.Example) (map[string]types.AttributeValue, error) {
	return attributevalue.MarshalMap(e)
}

func ExampleDynamoToModel(e map[string]types.AttributeValue) model.Example {
	var example model.Example

	attributevalue.UnmarshalMap(e, &example)

	return example
}
