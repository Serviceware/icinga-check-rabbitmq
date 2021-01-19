package internal

import rabbithole "github.com/Serviceware/rabbit-hole/v2"

// Checks if all connections are in running state
func CheckConnections(client *rabbithole.Client) int {
	connections, err := client.ListConnections()

	if err != nil {
		println(err.Error())
		return UNKNOWN
	}

	status := OK

	for _, connection := range connections {
		if connection.State != "running" {
			println(connection.Name + ".state=" + connection.State)
			status = WARNING
		}
	}

	return status
}
