package conf

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"os"
	"strconv"
)

type Specification struct {
	Port        int64          `default:"5432" required:"true"`
	DbUrl       string         `default:"postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable" required:"true"`
	Users       []string       `default:"" required:"false"`
	SentryUrl   string         `required:"true"`
	KafkaBroker map[string]int `required:"false"`
	Id          string         `required:"false"`
}

func SetEnv() (error, *Specification) {
	os.Setenv("MYAPP_USERS", "Roman,Andrey,Dima")
	os.Setenv("MYAPP_PORT", "5432")
	if _, err := strconv.Atoi(os.Getenv("MYAPP_PORT")); err != nil {
		return fmt.Errorf("Port must be type of int"), nil
	}
	os.Setenv("MYAPP_DBRL", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable")
	os.Setenv("MYAPP_SENTRYURL", "http://sentry:9000")
	if os.Getenv("MYAPP_SENTRYURL") == "" {
		return fmt.Errorf("Sentry url cannot be empty"), nil
	}
	os.Setenv("MYAPP_KAFKABROKER", "kafka:9092")
	os.Setenv("MYAPP_ID", "testid")
	conf := Specification{}
	err := envconfig.Process("myapp", &conf)
	if err != nil {
		return fmt.Errorf("can't process the config: %w", err), nil
	}
	return nil, &conf
}
