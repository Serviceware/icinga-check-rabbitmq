package test

import (
	"bitbucket.org/sabio-it/icinga-check-rabbitmq/checks"
	"testing"
)

func TestChannels(t *testing.T) {
	status := checks.CheckQueues(client())

	if status != 0 {
		t.Fail()
	}
}

func TestConnections(t *testing.T) {
	status := checks.CheckConnections(client())

	if status != 0 {
		t.Fail()
	}
}

func TestMessages(t *testing.T) {
	opts := &checks.CheckMessagesOpts{
		WarnLimit:     10,
		CriticalLimit: 20,
	}
	status := checks.CheckMessages(client(), opts)

	if status != 0 {
		t.Fail()
	}
}

// test is disabled atm because I don't know how to configure hostname for rabbitmq service in github actions
//func TestNode(t *testing.T) {
//	opts := &checks.CheckNodeOpts{
//		Node: "rabbit@github",
//	}
//	status := checks.CheckNode(client(), opts)
//
//	if status != 0 {
//		t.Fail()
//	}
//}

func TestPing(t *testing.T) {
	status := checks.Ping(client())

	if status != 0 {
		t.Fail()
	}
}

func TestQueues(t *testing.T) {
	status := checks.CheckQueues(client())

	if status != 0 {
		t.Fail()
	}
}
