package main

import (
	"os"
)

type Specification struct {
	Port         int `default:"5432"`
	Db_url       string
	Users        []string
	Sentry_url   string
	Kafka_broker map[string]int `default:"kafka:9092"`
	Id           string
}

func SetEnv() (users, port, db_url, sentry_url, kafka_broker, id string) {
	os.Setenv("MYAPP_USERS", "Roman,Andrey,Dima")
	os.Setenv("MYAPP_PORT", "5432")
	os.Setenv("MYAPP_DB_URL", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable")
	os.Setenv("MYAPP_SENTRY_URL", "http://sentry:9000")
	os.Setenv("MYAPP_KAFKA_BROKER", "kafka:9092")
	os.Setenv("MYAPP_ID", "testid")
	users = os.Getenv("MYAPP_USERS")
	port = os.Getenv("MYAPP_PORT")
	db_url = os.Getenv("MYAPP_DB_URL")
	sentry_url = os.Getenv("MYAPP_SENTRY_URL")
	kafka_broker = os.Getenv("MYAPP_KAFKA_BROKER")
	id = os.Getenv("MYAPP_ID")
	return 
}