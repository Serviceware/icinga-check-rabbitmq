package test

import (
	"bitbucket.org/sabio-it/icinga-check-rabbitmq/internal"
	"testing"
)

func TestHealthAlarms(t *testing.T) {
	status := internal.CheckHealth(client(), internal.ALARMS, &internal.CheckHealthOpts{})

	if status != 0 {
		t.Fail()
	}
}

func TestHealthLocalAlarms(t *testing.T) {
	status := internal.CheckHealth(client(), internal.LOCAL_ALARMS, &internal.CheckHealthOpts{})

	if status != 0 {
		t.Fail()
	}
}
