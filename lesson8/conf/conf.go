package conf

import (
	"os"
)

type Specification struct {
	Port         int `default:"5432"`
	DbUrl       string
	Users        []string
	SentryUrl   string
	KafkaBroker map[string]int `default:"kafka:9092"`
	Id           string
}

func SetEnv() (users, port, dbUrl, sentryUrl, kafkaBroker, id string) {
	os.Setenv("MYAPP_USERS", "Roman,Andrey,Dima")
	os.Setenv("MYAPP_PORT", "5432")
	os.Setenv("MYAPP_DBURL", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable")
	os.Setenv("MYAPP_SENTRYURL", "http://sentry:9000")
	os.Setenv("MYAPP_KAFKABROKER", "kafka:9092")
	os.Setenv("MYAPP_ID", "testid")
	users = os.Getenv("MYAPP_USERS")
	port = os.Getenv("MYAPP_PORT")
	dbUrl = os.Getenv("MYAPP_DB_URL")
	sentryUrl = os.Getenv("MYAPP_SENTRY_URL")
	kafkaBroker = os.Getenv("MYAPP_KAFKA_BROKER")
	id = os.Getenv("MYAPP_ID")
	return 
}