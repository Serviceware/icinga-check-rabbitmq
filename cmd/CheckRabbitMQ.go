package main

import (
	"bitbucket.org/sabio-it/icinga-check-rabbitmq/checks"
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
	PasswordFile string `long:"password-file" description:"File which contains the password for RabbitMQ authentication" group:"connection"`

	Ping        checks.Void              `command:"ping"`
	Health      checks.CheckHealthOpts   `command:"health"`
	Node        checks.CheckNodeOpts     `command:"node"`
	Messages    checks.CheckMessagesOpts `command:"messages"`
	Channels    checks.Void              `command:"channels"`
	Connections checks.Void              `command:"connections"`
	Queues      checks.Void              `command:"queues"`
}

var opts = new(args)

func main() {
	parser := parseFlags()

	status := 4
	switch parser.Active.Name {
	case "channels":
		status = checks.CheckChannels(rabbitmqClient())
	case "connections":
		status = checks.CheckConnections(rabbitmqClient())
	case "health":
		status = checks.CheckHealth(rabbitmqClient(), checks.Check(parser.Active.Active.Name), &opts.Health)
	case "messages":
		status = checks.CheckMessages(rabbitmqClient(), &opts.Messages)
	case "node":
		status = checks.CheckNode(rabbitmqClient(), &opts.Node)
	case "ping":
		status = checks.Ping(rabbitmqClient())
	case "queues":
		status = checks.CheckQueues(rabbitmqClient())
	}

	os.Exit(status)
}

func rabbitmqClient() *rabbithole.Client {
	password := readPassword()
	config := checks.CLientConfig{
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
