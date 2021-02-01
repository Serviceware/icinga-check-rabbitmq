package test

import (
	"bitbucket.org/sabio-it/icinga-check-rabbitmq/checks"
	"testing"
)

func TestHealthAlarms(t *testing.T) {
	status := checks.CheckHealth(client(), checks.ALARMS, &checks.CheckHealthOpts{})

	if status != 0 {
		t.Fail()
	}
}

func TestHealthLocalAlarms(t *testing.T) {
	status := checks.CheckHealth(client(), checks.LOCAL_ALARMS, &checks.CheckHealthOpts{})

	if status != 0 {
		t.Fail()
	}
}

func TestHealthProtocol(t *testing.T) {
	opts := &checks.CheckHealthOpts{
		ProtocolLister: checks.ProtocolListenerOpts{
			Protocol: "amqp",
		},
	}
	status := checks.CheckHealth(client(), checks.PROTOCOL_LISTENER, opts)

	println(status)
}

func TestHealthProtocolUnknown(t *testing.T) {
	opts := &checks.CheckHealthOpts{
		ProtocolLister: checks.ProtocolListenerOpts{
			Protocol: "123",
		},
	}
	status := checks.CheckHealth(client(), checks.PROTOCOL_LISTENER, opts)

	println(status)
}

func TestHealthPort(t *testing.T) {
	opts := &checks.CheckHealthOpts{
		PortListener: checks.PortListenerOpts{
			Port: 5672,
		},
	}
	status := checks.CheckHealth(client(), checks.PORT_LISTENER, opts)

	println(status)
}

func TestHealthPortUnknown(t *testing.T) {
	opts := &checks.CheckHealthOpts{
		PortListener: checks.PortListenerOpts{
			Port: 5671,
		},
	}
	status := checks.CheckHealth(client(), checks.PORT_LISTENER, opts)

	println(status)
}
