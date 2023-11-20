package exercise

import (
	"context"
	"github.com/google/wire"
	"github.com/gym-api/internal/platform/dynamodb"
)

type (
	creator interface {
		Set(ctx context.Context, e Exercise) error
	}
	cache interface {
		creator
	}
)

type serviceImpl struct {
	cache    cache
	dynamoDB dynamodb.DynamoDB[Entity]
}

func newService(cache cache, dynamoDB dynamodb.DynamoDB[Entity]) serviceImpl {
	return serviceImpl{
		cache:    cache,
		dynamoDB: dynamoDB,
	}
}

func (s serviceImpl) Add(ctx context.Context, e Exercise) (Exercise, error) {
	//err := s.cache.Set(ctx, e)
	//if err != nil {
	//	return Exercise{}, err
	//}
	_, err := s.dynamoDB.Set(ctx, NewEntity(e))
	if err != nil {
		return Exercise{}, err
	}
	return e, nil
}

var providerService = wire.NewSet(
	providerRedisCache,
	providerDynamoDB,
	newService,

	wire.Bind(new(cache), new(redisCache)),
)
