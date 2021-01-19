package test

import (
	"bitbucket.org/sabio-it/icinga-check-rabbitmq/internal"
	rabbithole "github.com/Serviceware/rabbit-hole/v2"
	"log"
)

func client() *rabbithole.Client {
	config := internal.CLientConfig{
		Address:  "http://localhost:15672",
		Username: "monitoring",
		Password: "secret",
	}

	client, err := config.NewRabbitMQClient()

	if err != nil {
		log.Fatal(err.Error())
	}

	return client
}
