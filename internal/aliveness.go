package internal

import (
	rabbithole "github.com/Serviceware/rabbit-hole/v2"
)

type AlivenessCheck struct {
	client *rabbithole.Client
}

type Aliveness struct {
	Status string `json:"status"`
}

func NewAlivenessCheck(client *rabbithole.Client) Check {
	return &AlivenessCheck{client: client}
}

func (c *AlivenessCheck) DoCheck() int {
	aliveness, err := c.client.Aliveness("labs")

	if err != nil {
		println(err.Error())
		return 2
	}

	println(aliveness.Status)
	return 0
}
