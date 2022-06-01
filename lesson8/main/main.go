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

	var s conf.Specification
	if err := conf.SetEnv(); err != nil {
		panic(err)
	}

	err := envconfig.Process("myapp", &s)
	if err != nil {
		log.Fatal(err.Error())
	}

	format := "Port: %v\nDb url: %v\nSentry url: %s\nId: %s\nKafka broker: %v\n"
	_, err = fmt.Printf(format, s.DbUrl, s.Port, s.SentryUrl, s.Id, s.KafkaBroker)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Users:")
	for _, u := range s.Users {
		fmt.Printf("  %s\n", u)
	}
}
