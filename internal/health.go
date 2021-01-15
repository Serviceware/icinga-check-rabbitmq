package internal

import rabbithole "github.com/Serviceware/rabbit-hole/v2"

type HealthCheck struct {
	client *rabbithole.Client
}

func NewHealthCheck(client *rabbithole.Client) Check {
	return &HealthCheck{client: client}
}

func (c *HealthCheck) DoCheck() int {
	health, err := c.client.HealthCheck()

	if err != nil {
		println(err.Error())
		return 2
	}

	if health.Status != "ok" {
		println("status =", health.Status)
		println("reason =", health.Reason)
		return 2
	} else {
		println("ok")
	}

	return 0
}
