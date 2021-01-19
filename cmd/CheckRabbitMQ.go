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
	Address string `long:"address" description:"Address of the nomad server" group:"connection" required:"true"`

	CaCert       string `long:"ca" description:"Path to ca cert" group:"connection"`
	ClientCert   string `long:"cert" description:"Path to client cert" group:"connection"`
	ClientKey    string `long:"key" description:"Path to client key" group:"connection"`
	Username     string `long:"username" description:"Username for RabbitMQ authentication" group:"connection"`
	Password     string `long:"password" description:"Password for RabbitMQ authentication" group:"connection"`
	PasswordFile string `long:"passwordFile" description:"File which contains the password for RabbitMQ authentication" group:"connection"`

	Ping        internal.Void              `command:"ping"`
	Health      internal.CheckHealthOpts   `command:"health"`
	Node        internal.CheckNodeOpts     `command:"node"`
	Messages    internal.CheckMessagesOpts `command:"messages"`
	Channels    internal.Void              `command:"channels"`
	Connections internal.Void              `command:"connections"`
	Queues      internal.Void              `command:"queues"`
}

var opts = new(args)

func main() {
	parser := parseFlags()

	status := 4
	switch parser.Active.Name {
	case "ping":
		status = internal.Ping(rabbitmqClient())
	case "health":
		status = internal.CheckHealth(rabbitmqClient(), internal.Check(parser.Active.Active.Name), &opts.Health)
	case "node":
		status = internal.CheckNode(rabbitmqClient(), &opts.Node)
	case "messages":
		status = internal.CheckMessages(rabbitmqClient(), &opts.Messages)
	case "connections":
		status = internal.CheckConnections(rabbitmqClient())
	case "queues":
		status = internal.CheckQueues(rabbitmqClient())
	}

	os.Exit(status)
}

func rabbitmqClient() *rabbithole.Client {
	password := readPassword()
	config := internal.CLientConfig{
		Address:    opts.Address,
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
