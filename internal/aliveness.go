package internal

import rabbithole "github.com/michaelklishin/rabbit-hole/v2"

type AlivenessCheck struct {
	client *rabbithole.Client
}

func NewAlivenessCheck(client *rabbithole.Client) Check {
	return &AlivenessCheck{client: client}
}

func (c *AlivenessCheck) DoCheck() int {
	return 0
}
