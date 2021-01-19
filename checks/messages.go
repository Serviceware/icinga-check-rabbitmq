package checks

import rabbithole "github.com/Serviceware/rabbit-hole/v2"

type CheckMessagesOpts struct {
	WarningLimit  int `long:"total-messages-warning-limit" description:""`
	CriticalLimit int `long:"total-messages-critical-limit" description:""`
}

// Checks if message count is above a warn or critical limit
func CheckMessages(client *rabbithole.Client, opts *CheckMessagesOpts) int {
	overview, err := client.Overview()

	if err != nil {
		println(err.Error())
		return UNKNOWN
	}

	status := OK

	totalMessages := overview.QueueTotals.Messages
	if totalMessages > opts.CriticalLimit {
		println(overview.QueueTotals.Messages, "messages exceeds critical limit", opts.CriticalLimit)
		status = CRITICAL
	} else if totalMessages > opts.WarningLimit {
		println(overview.QueueTotals.Messages, "messages exceeds warn limit", opts.WarningLimit)
		status = WARNING
	}

	if status == OK {
		println("ok")
	}

	return status
}
