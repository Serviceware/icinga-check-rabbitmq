package internal

import (
	rabbithole "github.com/Serviceware/rabbit-hole/v2"
)

func CheckAliveness(client *rabbithole.Client, vhost string) int {
	aliveness, err := client.Aliveness(vhost)

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
