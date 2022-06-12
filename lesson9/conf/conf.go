package conf

import (
	"encoding/json"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"os"
)

type Specification struct {
	Port        string   `json:"port"`
	DbUrl       string   `json:"dbUrl"`
	Users       []string `json:"users"`
	SentryUrl   string   `json:"sentryUrl"`
	KafkaBroker string   `json:"kafkaBroker"`
	Id          string   `json:"id"`
}

func SetEnv() (error, *Specification) {
	confJson, err := os.Open("../conf/conf.json")
	if err != nil {
		return fmt.Errorf("Cannot open file: %v", err), nil
	}
	defer func() error {
		err := confJson.Close()
		if err != nil {
			return fmt.Errorf("Cannot close file: %v", err)
		}
		return nil
	}()
	config := Specification{}
	err = json.NewDecoder(confJson).Decode(&config)
	if err != nil {
		return fmt.Errorf("Cannot decode json file into structure: %v", err), nil
	}
	err = envconfig.Process("myapp", config)
	if err != nil {
		return err, nil
	}
	return nil, &config
}
