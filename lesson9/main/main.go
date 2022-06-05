package main

import (
	"Go_level_one/lesson8/conf"
	"fmt"
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

	format := "Port: %v\nDb url: %v\nSentry url: %v\nId: %s\n"
	_, err = fmt.Printf(format, config.DbUrl, config.Port, config.SentryUrl, config.Id)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Users:")
	for _, u := range config.Users {
		fmt.Printf("  %v\n", u)
	}

	fmt.Println("Kafka broker:")
	for _, u := range config.KafkaBroker {
		fmt.Printf("  %v\n", u)
	}
}
