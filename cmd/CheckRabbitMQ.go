package main

import (
	"bitbucket.org/sabio-it/icinga-check-rabbitmq/internal"
	"github.com/jessevdk/go-flags"
	rabbithole "github.com/michaelklishin/rabbit-hole/v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type args struct {
	Address string `long:"address" description:"Address of the nomad server" group:"connection"`

	CaCert          string `long:"ca" description:"Path to ca cert" group:"connection"`
	ClientCert      string `long:"cert" description:"Path to client cert" group:"connection"`
	ClientKey       string `long:"key" description:"Path to client key" group:"connection"`
	Username        string `long:"username" description:"Username for RabbitMQ authentication" group:"connection"`
	Password        string `long:"password" description:"Password for RabbitMQ authentication" group:"connection"`
	CredentialsFile string `long:"credentialsFile" description:"Path to a file which contains credentials" group:"connection"`

	Aliveness void `command:"aliveness"`
	Check     void `command:"check"`
}

type void struct {
}

type credentials struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var opts = new(args)

func main() {
	parser := parseFlags()

	code := 4
	switch parser.Active.Name {
	case "aliveness":
		code = internal.NewClusterCheck(rabbitmqClient()).DoCheck()
	}

	os.Exit(code)
}

func rabbitmqClient() *rabbithole.Client {
	credentials := readCredentials()
	config := internal.CLientConfig{
		CaCert:     opts.CaCert,
		ClientCert: opts.ClientCert,
		ClientKey:  opts.ClientKey,
		Username:   credentials.Username,
		Password:   credentials.Password,
	}
	client, err := config.NewRabbitMQClient()

	if err != nil {
		log.Fatal(err.Error())
	}

	return client
}

func readCredentials() *credentials {
	if opts.CredentialsFile != "" {
		credentials := readCredentialsFromFile()

		if credentials != nil {
			return credentials
		}
	}

	return &credentials{
		Username: opts.Username,
		Password: opts.Password,
	}
}

func readCredentialsFromFile() *credentials {
	data, err := ioutil.ReadFile(opts.CredentialsFile)

	if err != nil {
		println("cannot read file", opts.CredentialsFile, ":", err.Error())
		return nil
	}

	credentials := &credentials{}
	err = yaml.Unmarshal(data, credentials)

	if err != nil {
		println("cannot read credentials from", opts.CredentialsFile, ":", err.Error())
		return nil
	}

	return credentials
}

func parseFlags() *flags.Parser {
	p := flags.NewParser(opts, flags.Default)
	_, err := p.Parse()

	if err != nil {
		log.Fatal(err.Error())
	}

	return p
}
