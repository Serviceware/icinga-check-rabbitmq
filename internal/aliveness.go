package internal

import (
	rabbithole "github.com/Serviceware/rabbit-hole/v2"
)

type CheckAlivenessOpts struct {
	Vhost string `long:"vhost" description:"The vhost to check"`
}

func CheckAliveness(client *rabbithole.Client, opts *CheckAlivenessOpts) int {
	aliveness, err := client.Aliveness(opts.Vhost)

	if err != nil {
		println(err.Error())
		return CRITICAL
	}

	println(aliveness.Status)

	if aliveness.Status != "ok" {
		return WARNING
	}

	return OK
}
