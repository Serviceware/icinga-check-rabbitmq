package internal

import (
	rabbithole "github.com/Serviceware/rabbit-hole/v2"
)

type AlivenessCheck struct {
	client *rabbithole.Client
	vhost  string
}

type Aliveness struct {
	Status string `json:"status"`
}

func NewAlivenessCheck(client *rabbithole.Client, vhost string) Check {
	println(client.Endpoint)
	return &AlivenessCheck{client: client, vhost: vhost}
}

func (c *AlivenessCheck) DoCheck() int {
	println(c.client.Endpoint)
	aliveness, err := c.client.Aliveness(c.vhost)

	if err != nil {
		println(err.Error())
		return 2
	}

	println(aliveness.Status)
	return 0
}
