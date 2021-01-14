package test

import (
	"bitbucket.org/sabio-it/icinga-check-rabbitmq/internal"
	rabbithole "github.com/Serviceware/rabbit-hole/v2"
	"log"
	"os"
	"testing"
)

func legacyClient() *rabbithole.Client {
	config := internal.CLientConfig{
		Address:    "https://rabbitmq01.nomadsupport-internal.hc.sabio.de:15671",
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

func TestLegacyAliveness(t *testing.T) {
	status := internal.NewAlivenessCheck(legacyClient()).DoCheck()

	if status != 0 {
		t.Fail()
	}
}
