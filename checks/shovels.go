package checks

import rabbithole "github.com/Serviceware/rabbit-hole/v2"

type CheckShovelsOpts struct {
	Vhost string `long:"vhost" description:"vhost to check"`
}

// Checks if all queues are in state running
func CheckShovels(client *rabbithole.Client, opts *CheckShovelsOpts) int {
	shovels, err := client.ListShovelStatus(opts.Vhost)

	if err != nil {
		println(err.Error())
		return UNKNOWN
	}

	status := OK

	for _, shovel := range shovels {
		if shovel.State != "running" || shovel.State != "starting" {
			println(shovel.Name + ".state=" + shovel.State)
			status = WARNING
		}
	}

	if status == OK {
		println("ok - all shovels in state running or starting ")
	}

	return status
}
