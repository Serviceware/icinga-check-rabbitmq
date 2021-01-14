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
	println(1)
	aliveness, err := c.client.Aliveness("labs")
	println(2)
	if err != nil {
		println(err.Error())
		return 2
	}
	println(3)

	println(aliveness.Status)
	println(4)
	return 0
}
