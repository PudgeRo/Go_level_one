package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
)

func main() {
    
	var s Specification
    SetEnv()
	err := envconfig.Process("myapp", &s)
	if err != nil {
		log.Fatal(err.Error())
	}

    os.Setenv("MYAPP_SENTRY_URL", "http://sentry:9000")

	format := "Port: %v\nDb_url: %v\nSentry_url: %s\nId: %s\n"
	_, err = fmt.Printf(format, s.Db_url, s.Port, s.Sentry_url, s.Id)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Users:")
	for _, u := range s.Users {
		fmt.Printf("  %s\n", u)
	}

	fmt.Println("Kafka broker:")
	for k, v := range s.Kafka_broker {
		fmt.Printf("  %s: %d\n", k, v)
	}
}
