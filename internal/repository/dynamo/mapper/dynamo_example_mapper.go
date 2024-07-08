package mapper

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/charmingruby/tynamo/internal/model"
)

func ExampleToDynamo(e *model.Example) (map[string]types.AttributeValue, error) {
	return attributevalue.MarshalMap(e)
}
