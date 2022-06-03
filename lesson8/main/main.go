package main

import (
	"Go_level_one/lesson8/conf"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()

	err, config := conf.SetEnv()
	if err != nil {
		panic(err)
	}

	err = envconfig.Process("myapp", &config)
	if err != nil {
		log.Fatal(err.Error())
	}

	format := "Port: %v\nDb url: %v\nSentry url: %v\nId: %s\nKafka broker: %v\n"
	_, err = fmt.Printf(format, config.DbUrl, config.Port, config.SentryUrl, config.Id, config.KafkaBroker)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Users:")
	for _, u := range config.Users {
		fmt.Printf("  %v\n", u)
	}
}
