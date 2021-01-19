package test

import (
	"bitbucket.org/sabio-it/icinga-check-rabbitmq/internal"
	"testing"
)

func TestChannels(t *testing.T) {
	status := internal.CheckQueues(client())

	if status != 0 {
		t.Fail()
	}
}

func TestConnections(t *testing.T) {
	status := internal.CheckConnections(client())

	if status != 0 {
		t.Fail()
	}
}

func TestMessages(t *testing.T) {
	opts := &internal.CheckMessagesOpts{
		WarnLimit:     10,
		CriticalLimit: 20,
	}
	status := internal.CheckMessages(client(), opts)

	if status != 0 {
		t.Fail()
	}
}

func TestNode(t *testing.T) {
	opts := &internal.CheckNodeOpts{
		Node: "rabbit@github",
	}
	status := internal.CheckNode(client(), opts)

	if status != 0 {
		t.Fail()
	}
}

func TestPing(t *testing.T) {
	status := internal.Ping(client())

	if status != 0 {
		t.Fail()
	}
}

func TestQueues(t *testing.T) {
	status := internal.CheckQueues(client())

	if status != 0 {
		t.Fail()
	}
}
