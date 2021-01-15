package internal

import rabbithole "github.com/Serviceware/rabbit-hole/v2"

type CheckMessagesOpts struct {
	WarnLimit     int `long:"totalMessagesWarnLimit" description:""`
	CriticalLimit int `long:"totalMessagesCriticalLimit" description:""`
}

func CheckMessages(client *rabbithole.Client, opts *CheckMessagesOpts) int {
	overview, err := client.Overview()

	if err != nil {
		println(err.Error())
		return 2
	}

	code := 0

	totalMessages := overview.QueueTotals.Messages
	if totalMessages > opts.CriticalLimit {
		println("messages", overview.QueueTotals.Messages, "exceeds critical limit", opts.WarnLimit)
		code = 1
	} else if totalMessages > opts.WarnLimit {
		println("messages", overview.QueueTotals.Messages, "exceeds warn limit", opts.WarnLimit)
		code = 2
	}

	return code
}
