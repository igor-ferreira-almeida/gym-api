package exercise

import (
	"github.com/google/wire"
	"github.com/gym-api/internal/platform/config"
	"github.com/gym-api/internal/platform/dynamodb"
	"strings"
)

type Entity struct {
	ID          string `dynamodbav:"id"`
	Name        string `dynamodbav:"name"`
	MuscleGroup string `dynamodbav:"muscle_group"`
}

func (e Entity) GetPartitionKey() string {
	return e.ID
}

func NewEntity(e Exercise) Entity {
	return Entity{
		ID:          strings.ToLower(strings.ReplaceAll(e.Name, " ", "_")),
		Name:        e.Name,
		MuscleGroup: e.MuscleGroup,
	}
}

func newDynamoDB() dynamodb.DynamoDB[Entity] {
	isLocal = true
	cfg := config.DynamoDBConfig{
		TableName:       "exercise",
		Region:          "us-west-1",
		AccessKeyID:     "dummy",
		SecretAccessKey: "dummy",
		SessionToken:    "dummy",
		Endpoint:        "http://localhost",
		Port:            "8000",
	}
	return dynamodb.NewDynamoDB[Entity](isLocal, cfg)
}

var providerDynamoDB = wire.NewSet(
	newDynamoDB,
)
