package config

type DynamoDBConfig struct {
	TableName       string `json:"table_name"`
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	SessionToken    string `json:"session_token"`
	Region          string `json:"region"`
	Endpoint        string `json:"endpoint"`
	Port            string `json:"port"`
}
