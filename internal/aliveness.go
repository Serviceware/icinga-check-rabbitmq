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
	return &AlivenessCheck{client: client, vhost: vhost}
}

func (c *AlivenessCheck) DoCheck() int {
	aliveness, err := c.client.Aliveness(c.vhost)

	if err != nil {
		println(err.Error())
		return 2
	}

	println(aliveness.Status)

	if aliveness.Status != "ok" {
		return 1
	}

	return 0
}
