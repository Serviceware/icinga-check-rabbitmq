package main

import (
	"bitbucket.org/sabio-it/icinga-check-rabbitmq/internal"
	rabbithole "github.com/Serviceware/rabbit-hole/v2"
	"github.com/jessevdk/go-flags"
	"io/ioutil"
	"log"
	"os"
)

type args struct {
	Address string `long:"address" description:"Address of the nomad server" group:"connection"`

	CaCert       string `long:"ca" description:"Path to ca cert" group:"connection"`
	ClientCert   string `long:"cert" description:"Path to client cert" group:"connection"`
	ClientKey    string `long:"key" description:"Path to client key" group:"connection"`
	Username     string `long:"username" description:"Username for RabbitMQ authentication" group:"connection"`
	Password     string `long:"password" description:"Password for RabbitMQ authentication" group:"connection"`
	PasswordFile string `long:"passwordFile" description:"File which contains the password for RabbitMQ authentication" group:"connection"`

	Aliveness void `command:"aliveness"`
	Check     void `command:"check"`
}

type void struct {
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
	password := readPassword()
	config := internal.CLientConfig{
		CaCert:     opts.CaCert,
		ClientCert: opts.ClientCert,
		ClientKey:  opts.ClientKey,
		Username:   opts.Username,
		Password:   password,
	}
	client, err := config.NewRabbitMQClient()

	if err != nil {
		log.Fatal(err.Error())
	}

	return client
}

func readPassword() string {
	if opts.PasswordFile != "" {
		credentials := readCredentialsFromFile()

		if credentials != "" {
			return credentials
		}
	}

	return opts.Password
}

func readCredentialsFromFile() string {
	data, err := ioutil.ReadFile(opts.PasswordFile)

	if err != nil {
		println("cannot read file", opts.PasswordFile, ":", err.Error())
		return ""
	}

	return string(data)
}

func parseFlags() *flags.Parser {
	p := flags.NewParser(opts, flags.Default)
	_, err := p.Parse()

	if err != nil {
		log.Fatal(err.Error())
	}

	return p
}
