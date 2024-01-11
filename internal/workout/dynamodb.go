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
	Name        string      `dynamodbav:"name"`
	MuscleGroup string      `dynamodbav:"muscle_group"`
	Sets        []SetEntity `dynamodbav:"sets"`
}

type SetEntity struct {
	Reps               uint          `dynamodbav:"reps"`
	Weight             uint          `dynamodbav:"weight"`
	Rest               time.Duration `dynamodbav:"rest"`
	GravitationalForce float64       `dynamodbav:"gravitational_force"`
}

func (e Entity) GetPartitionKey() string {
	return e.ID
}

func (e Entity) ToWorkout() (Workout, error) {
	date, err := time.Parse("2006-01-02", e.ID)
	exercises := make([]Exercise, len(e.Exercises))
	for i, ex := range e.Exercises {
		exercises[i].Name = ex.Name
		exercises[i].MuscleGroup = ex.MuscleGroup
		sets := make([]Set, len(ex.Sets))
		for si, s := range ex.Sets {
			sets[si].Reps = s.Reps
			sets[si].Weight = s.Weight
			sets[si].Rest = s.Rest
			sets[si].GravitationalForce = s.GravitationalForce
		}
		exercises[i].Sets = sets
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
		sets := make([]SetEntity, len(e.Sets))
		for si, s := range e.Sets {
			sets[si].Reps = s.Reps
			sets[si].Weight = s.Weight
			sets[si].Rest = s.Rest
			sets[si].GravitationalForce = s.GravitationalForce
		}
		exercises[i].Sets = sets
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
