package internal

import rabbithole "github.com/Serviceware/rabbit-hole/v2"

// Checks if there is a blocked channel
func CheckChannel(client *rabbithole.Client) int {
	channels, err := client.ListChannels()

	if err != nil {
		println(err.Error())
		return UNKNOWN
	}

	status := OK

	for _, channel := range channels {
		if channel.ClientFlowBlocked {
			println(channel.Name + " is client flow blocked")
			status = WARNING
		}
	}

	return status
}
