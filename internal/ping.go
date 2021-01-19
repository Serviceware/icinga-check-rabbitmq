package internal

import (
	rabbithole "github.com/Serviceware/rabbit-hole/v2"
)

func Ping(client *rabbithole.Client) int {
	_, err := client.Whoami()

	if err != nil {
		println(err.Error())
		return CRITICAL
	}

	println("ok - node is reachable")
	return OK
}
