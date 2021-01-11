package test

import (
	"bitbucket.org/sabio-it/icinga-check-rabbitmq/internal"
	rabbithole "github.com/michaelklishin/rabbit-hole/v2"
	"log"
	"os"
	"testing"
)

func legacyClient() *rabbithole.Client {
	config := internal.CLientConfig{
		Address:    "https://infrastructure-rabbitmq-mgmt.service.fsn.consul-internal.sabio.de:15671",
		CaCert:     "legacy/ca.pem",
		ClientCert: "legacy/cert.pem",
		ClientKey:  "legacy/key.pem",
		Username:   os.Getenv("RABBITMQ_USERNAME"),
		Password:   os.Getenv("RABBITMQ_PASSWORD"),
	}

	client, err := config.NewRabbitMQClient()

	if err != nil {
		log.Fatal(err.Error())
	}

	return client
}

func TestLegacyAlivenews(t *testing.T) {
	status := internal.NewAlivenessCheck(legacyClient()).DoCheck()

	if status != 0 {
		t.Fail()
	}
}
