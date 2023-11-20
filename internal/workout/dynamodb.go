package workout

import (
	"github.com/google/wire"
	"github.com/gym-api/internal/platform/config"
	"github.com/gym-api/internal/platform/dynamodb"
	"time"
)

type Entity struct {
	ID        string           `dynamodbav:"id"`
	Exercises []ExerciseEntity `dynamodbav:"exercises"`
}

type ExerciseEntity struct {
	Name        string `dynamodbav:"name"`
	MuscleGroup string `dynamodbav:"muscle_group"`
}

func (e Entity) GetPartitionKey() string {
	return e.ID
}

func (e Entity) ToWorkout() (Workout, error) {
	date, err := time.Parse("2006-01-02", e.ID)
	exercises := make([]Exercise, len(e.Exercises))
	for index, exercise := range e.Exercises {
		exercises[index].Name = exercise.Name
		exercises[index].MuscleGroup = exercise.MuscleGroup
	}
	return Workout{
		Date:      date,
		Exercises: exercises,
	}, err
}

func NewEntity(w Workout) Entity {
	exercises := make([]ExerciseEntity, len(w.Exercises))
	for i, e := range w.Exercises {
		exercises[i].Name = e.Name
		exercises[i].MuscleGroup = e.MuscleGroup
	}
	return Entity{
		ID:        w.Date.Format(time.DateOnly),
		Exercises: exercises,
	}
}

func newDynamoDB() dynamodb.DynamoDB[Entity] {
	cfg := config.DynamoDBConfig{
		TableName:       "workout",
		Region:          "us-west-1",
		AccessKeyID:     "dummy",
		SecretAccessKey: "dummy",
		SessionToken:    "dummy",
		Endpoint:        "http://localhost",
		Port:            "8000",
	}
	return dynamodb.NewDynamoDB[Entity](true, cfg)
}

var providerDynamoDB = wire.NewSet(
	newDynamoDB,
)
