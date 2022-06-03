package main

import (
	"Go_level_one/lesson8/conf"
	"fmt"
	"log"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	err, config := conf.SetEnv()
	if err != nil {
		panic(err)
	}
	format := "Port: %v\nDb url: %v\nSentry url: %v\nId: %v\n"
	_, err = fmt.Printf(format, config.DbUrl, config.Port, config.SentryUrl, config.Id)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Users:")
	for _, u := range config.Users {
		fmt.Printf("  %v\n", u)
	}

	fmt.Println("Kafka broker:")
	for k, v := range config.KafkaBroker {
		fmt.Printf("  %v: %v\n", k, v)
	}
}
