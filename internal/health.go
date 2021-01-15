package internal

import rabbithole "github.com/Serviceware/rabbit-hole/v2"

func CheckHealth(client *rabbithole.Client) int {
	health, err := client.HealthCheck()

	if err != nil {
		println(err.Error())
		return CRITICAL
	}

	if health.Status != "ok" {
		println("status =", health.Status)
		println("reason =", health.Reason)
		return WARNING
	} else {
		println("ok")
		return OK
	}
}
