package dynamodb

type Entity interface {
	GetPartitionKey() string // Used as key for dynamodb
}
