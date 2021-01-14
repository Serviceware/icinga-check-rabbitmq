package internal

import rabbithole "github.com/Serviceware/rabbit-hole/v2"

type ClusterCheck struct {
	client *rabbithole.Client
}

func NewClusterCheck(client *rabbithole.Client) Check {
	return &ClusterCheck{client: client}
}

func (c *ClusterCheck) DoCheck() int {
	return 0
}
