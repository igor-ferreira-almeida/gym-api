package workout

import (
	"context"
	"github.com/google/wire"
	"github.com/gym-api/internal/platform/dynamodb"
)

type (
	creator interface {
		Set(ctx context.Context, w Workout) error
	}
	cache interface {
		creator
	}
)

type serviceImpl struct {
	dynamoDB dynamodb.DynamoDB[Entity]
}

func newService(dynamoDB dynamodb.DynamoDB[Entity]) serviceImpl {
	return serviceImpl{
		dynamoDB: dynamoDB,
	}
}

func (s serviceImpl) Add(ctx context.Context, w Workout) (Workout, error) {
	_, err := s.dynamoDB.Set(ctx, NewEntity(w))
	if err != nil {
		return Workout{}, err
	}
	return w, nil
}

func (s serviceImpl) Get(ctx context.Context, id string) (Workout, error) {
	entity, err := s.dynamoDB.Get(ctx, id)
	if err != nil {
		return Workout{}, err
	}

	if err != nil {
		return Workout{}, err
	}
	return entity.ToWorkout()
}

var providerService = wire.NewSet(
	providerDynamoDB,
	newService,
)
