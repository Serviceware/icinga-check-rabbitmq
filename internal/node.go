package internal

import rabbithole "github.com/Serviceware/rabbit-hole/v2"

type CheckNodeOpts struct {
	Node string `long:"node" description:"The node which should be checked"`
}

func CheckNode(client *rabbithole.Client, opts *CheckNodeOpts) int {
	node, err := client.GetNode(opts.Node)

	if err != nil {
		println(err.Error())
		return 2
	}

	code := 0
	if node.DiskFreeAlarm {
		println("diskfree", node.DiskFree, "exceeds", node.DiskFreeLimit)
		code = 1
	}

	if node.MemAlarm {
		println("memfree", node.MemUsed, "exceeds", node.MemLimit)
		code = 1
	}

	if !node.IsRunning {
		println("node", opts.Node, "not running")
		code = 2
	}

	return code
}