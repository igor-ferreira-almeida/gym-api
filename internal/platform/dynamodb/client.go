package dynamodb

import (
	"context"
	"github.com/gym-api/internal/platform/config"
)

type DynamoDB[T Entity] interface {
	Get(ctx context.Context, key string) (T, error)
	Set(ctx context.Context, entity T) (T, error)
	Remove(ctx context.Context, key string) error
}

func NewDynamoDB[T Entity](isLocal bool, c config.DynamoDBConfig) DynamoDB[T] {
	return NewLocalDynamoDB[T](c)
}
