package test

import (
	"bitbucket.org/sabio-it/icinga-check-rabbitmq/internal"
	rabbithole "github.com/Serviceware/rabbit-hole/v2"
	"io/ioutil"
	"log"
	"testing"
)

func readCredentialsFromFile(passwordFile string) string {
	data, err := ioutil.ReadFile(passwordFile)

	if err != nil {
		println("cannot read file", passwordFile, ":", err.Error())
		return ""
	}

	println(string(data[:3]), "...", string(data[29:]))

	return string(data)
}

func legacyClient() *rabbithole.Client {
	config := internal.CLientConfig{
		Address:    "https://rabbitmq01.nomadsupport-internal.hc.sabio.de:15671",
		CaCert:     "legacy/ca.pem",
		ClientCert: "legacy/cert.pem",
		ClientKey:  "legacy/key.pem",
		Username:   "monitoring",
		Password:   readCredentialsFromFile("password"),
	}

	client, err := config.NewRabbitMQClient()

	if err != nil {
		log.Fatal(err.Error())
	}

	return client
}

func TestLegacyAliveness(t *testing.T) {
	status := internal.NewAlivenessCheck(legacyClient(), "production").DoCheck()

	if status != 0 {
		t.Fail()
	}
}

func TestLegacyHealth(t *testing.T) {
	status := internal.NewHealthCheck(legacyClient()).DoCheck()

	if status != 0 {
		t.Fail()
	}
}
