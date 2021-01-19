package internal

import rabbithole "github.com/Serviceware/rabbit-hole/v2"

type CheckMessagesOpts struct {
	WarnLimit     int `long:"totalMessagesWarnLimit" description:""`
	CriticalLimit int `long:"totalMessagesCriticalLimit" description:""`
}

// Checks if message count is above a warn or critical limit
func CheckMessages(client *rabbithole.Client, opts *CheckMessagesOpts) int {
	overview, err := client.Overview()

	if err != nil {
		println(err.Error())
		return UNKNOWN
	}

	code := OK

	totalMessages := overview.QueueTotals.Messages
	if totalMessages > opts.CriticalLimit {
		println(overview.QueueTotals.Messages, "messages exceeds critical limit", opts.CriticalLimit)
		code = CRITICAL
	} else if totalMessages > opts.WarnLimit {
		println(overview.QueueTotals.Messages, "messages exceeds warn limit", opts.WarnLimit)
		code = WARNING
	}

	return code
}
