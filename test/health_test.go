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
