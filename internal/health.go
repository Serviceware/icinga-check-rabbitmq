package internal

import rabbithole "github.com/Serviceware/rabbit-hole/v2"

func CheckHealth(client *rabbithole.Client) int {
	health, err := client.HealthCheck()

	if err != nil {
		println(err.Error())
		return 2
	}

	code := 0

	if health.Status != "ok" {
		println("status =", health.Status)
		println("reason =", health.Reason)
		code = 1
	} else {
		println("ok")
	}

	return code
}
