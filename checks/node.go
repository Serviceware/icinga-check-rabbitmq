package checks

import rabbithole "github.com/Serviceware/rabbit-hole/v2"

type CheckNodeOpts struct {
	Node string `long:"node" description:"The node which should be checked"`
}

// Checks if node is running and disk or memory alarms are raised
func CheckNode(client *rabbithole.Client, opts *CheckNodeOpts) int {
	node, err := client.GetNode(opts.Node)

	if err != nil {
		println(err.Error())
		return UNKNOWN
	}

	status := OK
	if node.DiskFreeAlarm {
		println("diskfree", node.DiskFree, "exceeds", node.DiskFreeLimit)
		status = WARNING
	}

	if node.MemAlarm {
		println("memfree", node.MemUsed, "exceeds", node.MemLimit)
		status = WARNING
	}

	if !node.IsRunning {
		println("node", opts.Node, "not running")
		status = CRITICAL
	}

	if status == OK {
		println("ok - node is running, no alarms")
	}

	return status
}
