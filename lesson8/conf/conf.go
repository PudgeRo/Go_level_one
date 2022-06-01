package conf

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

type Specification struct {
	Port        int `default: "4321"`
	DbUrl       string
	Users       []string
	SentryUrl   string
	KafkaBroker int `default: "9092"`
	Id          string
}

func SetEnv() error {
	conf, err := os.ReadFile("../conf/conf.json")
	if err != nil {
		return err
	}

	var configuration Specification
	if err := json.Unmarshal(conf, &configuration); err != nil {
		return err
	}

	os.Setenv("MYAPP_USERS", strings.Join(configuration.Users, ","))
	os.Setenv("MYAPP_PORT", strconv.Itoa(configuration.Port))
	os.Setenv("MYAPP_DBURL", configuration.DbUrl)
	os.Setenv("MYAPP_SENTRYURL", configuration.SentryUrl)
	os.Setenv("MYAPP_KAFKABROKER", strconv.Itoa(configuration.KafkaBroker))
	os.Setenv("MYAPP_ID", configuration.Id)
	return nil
}
