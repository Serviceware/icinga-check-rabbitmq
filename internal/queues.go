package internal

import rabbithole "github.com/Serviceware/rabbit-hole/v2"

// Checks if all queues are in state running
func CheckQueues(client *rabbithole.Client) int {
	queues, err := client.ListQueues()

	if err != nil {
		println(err.Error())
		return UNKNOWN
	}

	status := OK

	for _, queue := range queues {
		if queue.Status != "running" {
			println(queue.Name + ".state=" + queue.Status)
			status = WARNING
		}
	}

	if status == OK {
		println("ok")
	}

	return status
}
