//go:generate mockgen -source=local_dynamodb.go -destination=local_dynamodb_mock.go -package=kvs
package dynamodb

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gym-api/internal/platform/config"
	"log"
	"reflect"
)

var (
	LoadDefaultConfig = awscfg.LoadDefaultConfig
	NewFromConfig     = dynamodb.NewFromConfig
)

type LocalDynamoDB[T Entity] struct {
	dynamodbClient *dynamodb.Client
	tableName      *string
}

func NewLocalDynamoDB[T Entity](c config.DynamoDBConfig) LocalDynamoDB[T] {
	cfg, err := LoadDefaultConfig(
		context.TODO(),
		awscfg.WithRegion(c.Region),
		awscfg.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     c.AccessKeyID,
				SecretAccessKey: c.SecretAccessKey,
				SessionToken:    c.SessionToken,
				Source:          "Hard-coded credentials; values are irrelevant for local DynamoDB",
			},
		}),
		awscfg.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL: fmt.Sprintf("%s:%s", c.Endpoint, c.Port),
				}, nil
			})))
	if err != nil {
		panic(err)
	}
	client := NewFromConfig(cfg)

	return LocalDynamoDB[T]{
		dynamodbClient: client,
		tableName:      aws.String(c.TableName),
	}
}

func (client LocalDynamoDB[T]) Get(ctx context.Context, key string) (T, error) {
	var t T
	tag := reflect.ValueOf(&t).Elem().Type().Field(0).Tag.Get("dynamodbav")
	output, err := client.dynamodbClient.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: client.tableName,
		Key: map[string]types.AttributeValue{
			tag: &types.AttributeValueMemberS{Value: key},
		},
	})
	if err != nil {
		return *new(T), err
	}
	var entity T
	err = attributevalue.UnmarshalMap(output.Item, &entity)
	return entity, err
}

func (client LocalDynamoDB[T]) Set(ctx context.Context, entity T) (T, error) {
	attrValue, _ := attributevalue.MarshalMap(entity)
	output, err := client.dynamodbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: client.tableName,
		Item:      attrValue,
	})
	if err != nil {
		log.Fatal(err)
		return *new(T), err
	}

	err = attributevalue.UnmarshalMap(output.Attributes, &entity)

	return entity, err
}

func (k LocalDynamoDB[T]) Remove(ctx context.Context, key string) error {
	_, err := k.dynamodbClient.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: k.tableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: key},
		},
	})
	return err
}

//
//func (k LocalDynamoDB[T]) GetBatch(ctx context.Context, keys ...string) ([]dynamodb.EntityWrapper[T], error) {
//	var t T
//	tag := reflect.ValueOf(&t).Elem().Type().Field(0).Tag.Get("dynamodbav")
//	var keysAttr []map[string]types.AttributeValue
//
//	for _, key := range keys {
//		keysAttr = append(keysAttr, map[string]types.AttributeValue{
//			tag: &types.AttributeValueMemberS{Value: key},
//		})
//	}
//
//	output, err := k.client.BatchGetItem(ctx, &dynamodb.BatchGetItemInput{
//		RequestItems: map[string]types.KeysAndAttributes{
//			*k.tableName: {Keys: keysAttr},
//		},
//	})
//	if err != nil {
//		return []dynamodb.EntityWrapper[T]{}, err
//	}
//
//	var entities []dynamodb.EntityWrapper[T]
//	for _, res := range output.Responses[*k.tableName] {
//		var entity T
//		err = attributevalue.UnmarshalMap(res, &entity)
//		entities = append(entities, dynamodb.GetWrapper(dynamodb.WithEntity(entity)))
//	}
//
//	return entities, nil
//}
//
//func (k LocalDynamoDB[T]) GetBulk(ctx context.Context, keys ...string) ([]dynamodb.EntityWrapper[T], map[string]error, error) {
//	return []dynamodb.EntityWrapper[T]{}, nil, nil
//}
//

//func (k LocalDynamoDB[T]) SetBatch(
//	_ context.Context, _ []dynamodb.EntityWrapper[T], _ ...furykvs.OptionItem,
//) ([]dynamodb.EntityWrapper[T], error) {
//	return []dynamodb.EntityWrapper[T]{}, nil
//}
//
//func (k LocalDynamoDB[T]) SetBulk(
//	_ context.Context, _ []dynamodb.EntityWrapper[T], _ ...furykvs.OptionItem,
//) ([]dynamodb.EntityWrapper[T], map[string]error, error) {
//	return []dynamodb.EntityWrapper[T]{}, nil, nil
//}
//
//
//func (k LocalDynamoDB[T]) RemoveBatch(ctx context.Context, keys ...string) error {
//	return nil
//}
