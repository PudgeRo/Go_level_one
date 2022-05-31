package main

import (
	"Go_level_one/lesson8/conf"
	"fmt"
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
)

func main() {
	var s conf.Specification
	conf.SetEnv()
	err := envconfig.Process("myapp", &s)
	if err != nil {
		log.Fatal(err.Error())
	}

    os.Setenv("MYAPP_SENTRY_URL", "http://sentry:9000")

	format := "Port: %v\nDb url: %v\nSentry url: %s\nId: %s\n"
	_, err = fmt.Printf(format, s.DbUrl, s.Port, s.SentryUrl, s.Id)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Users:")
	for _, u := range s.Users {
		fmt.Printf("  %s\n", u)
	}

	fmt.Println("Kafka broker:")
	for k, v := range s.KafkaBroker {
		fmt.Printf("  %s: %d\n", k, v)
	}
}
