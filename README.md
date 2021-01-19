# Icinga check for RabbitMQ

Icinga check for RabbitMQ

## Build

```
go build -o check_rabbitmq ./cmd/CheckRabbitMQ.go
```

## Test

Start a docker container for RabbitMQ:
```
docker run -h github -e RABBITMQ_DEFAULT_USER=monitoring -e RABBITMQ_DEFAULT_PASS=secret -e RABBITMQ_NODENAME=rabbit -p 15672:15672 rabbitmq:3.8.10-rc.6-management
```

Run tests:
```
go test -v ./test
```

## Checks

### Ping

Just calls `/api/whoami` and return `CRITICAL` if any error occurs, otherwise `OK`.

```shell
check-rabbitmq ping
```


### Health

Calls the health checks `/api/health/check/*` (from version 3.8.10).

```shell
check-rabbitmq health alarms
```

```shell
check-rabbitmq health local-alarms
```

```shell
check-rabbitmq health certificate-expiration --within 1 --unit days
```

```shell
check-rabbitmq health port-listener --port 5671
```

```shell
check-rabbitmq health protocol-listener --protocol amqp
```

```shell
check-rabbitmq health node-is-mirror-sync-critical
```

```shell
check-rabbitmq health node-is-mirror-sync-critical
```

### Node

Checks if node is running and if a memory or disk alert is raised.

```shell
check-rabbitmq node --node rabbit
```


### Messages

Checks if the message count is above a warning or critical limit.

```shell
check-rabbitmq messages --total-messages-warning-limit 100000 --total-messages-critial-limit 200000
```


### Queues

Checks all queues if they are in `running` state. If one is not, a `WARNING` is returned

```shell
check-rabbitmq queues
```


### Channels

Checks all channels if they are client flow blocked. If one is not, a `WARNING` is returned.

```shell
check-rabbitmq channels
```


### Connections

Checks all connections if they are in state `running`. If one is not, a `WARNING` is returned. 

```shell
check-rabbtimq connections
```


## Configure Icinga

### CheckCommand

```
object CheckCommand "rabbitmq" {
    command = [ "/etc/icinga2/scripts/check_rabbitmq" ]

    arguments = {
        "check" = {
            value = "$rabbitmq_check$"
            skip_key = true
            order = -2
        }
        "health" = {
            value = "$rabbitmq_health_check$"
            skip_key = true
            order = -1
            set_if = {{ macro("$rabbitmq_check$") == "health" }}
        }
        "--address" = "$rabbitmq_address$"
        "--ca" = "$rabbitmq_ca$"
        "--cert" = "$rabbitmq_cert$"
        "--key" = "$rabbitmq_key$"
        "--username" = "monitoring"
        "--passwordFile" = "$rabbitmq_password_file$"
        "--vhost" = {
            value = "$rabbitmq_vhost$",
            set_if = {{ macro("$rabbitmq_check$") == "aliveness" }}
        }
        "--node" = {
            value = "$rabbitmq_node$"
            set_if = {{ macro("$rabbitmq_check$") == "node" }}
        }
        "--total-messages-warning-limit" = {
            value = "$rabbitmq_total_messages_warn_limit$"
            set_if = {{ macro("$rabbitmq_check$") == "messages" }}
        }
        "--total-messages-critical-limit" = {
            value = "$rabbitmq_total_messages_critical_limit$"
            set_if = {{ macro("$rabbitmq_check$") == "messages" }}
        }
        "--within" = {
            value = "$rabbitmq_health_cert_expires_within$"
            set_if = {{ (macro("$rabbitmq_check$") == "health") && (macro("$rabbitmq_health_check$") == "certificate-expiration") }}
        }
        "--unit" = {
            value = "$rabbitmq_health_cert_expires_unit$"
            set_if = {{ (macro("$rabbitmq_check$") == "health") && (macro("$rabbitmq_health_check$") == "certificate-expiration") }}
        }
        "--port" = {
            value = "$rabbitmq_health_port_listener_port$"
            set_if = {{ (macro("$rabbitmq_check$") == "health") && (macro("$rabbitmq_health_check$") == "port-listener") }}
        }
        "--protocol" = {
            value = "$rabbitmq_health_protocol_listener_protocol$"
            set_if = {{ (macro("$rabbitmq_check$") == "health") && (macro("$rabbitmq_health_check$") == "protocol-listener") }}
        }
    }

    vars.rabbitmq_address = "$rabbitmq_address$"
    vars.rabbitmq_ca = "$rabbitmq_ca$"
    vars.rabbitmq_cert = "$rabbitmq_cert$"
    vars.rabbitmq_key = "$rabbitmq_key$"
    vars.rabbitmq_password_file = "$rabbitmq_password_file$"

    vars.rabbitmq_check = "ping"

    vars.rabbitmq_vhost = "/"
    vars.rabbitmq_node = ""
    vars.rabbitmq_total_messages_warn_limit = 0
    vars.rabbitmq_total_messages_critical_limit = 0

    vars.rabbitmq_health_cert_expires_within = 1
    vars.rabbitmq_health_cert_expires_unit = "days"
    vars.rabbitmq_health_port_listener_port = "15672"
    vars.rabbitmq_health_protocol_listener_protocol = "amqp"
}
```

### Service

```
apply Service "rabbitmq-channels" {
    check_command = "rabbitmq"
    vars.rabbitmq_check = "channels"
    assign where host.vars.rabbitmq_server.active
}

apply Service "rabbitmq-connections" {
    check_command = "rabbitmq"
    vars.rabbitmq_check = "connections"
    assign where host.vars.rabbitmq_server.active
}

apply Service "rabbitmq-health" {
    check_command = "rabbitmq"
    vars.rabbitmq_check = "health"
    assign where host.vars.rabbitmq_server.active
}

apply Service "rabbitmq-messages" {
    check_command = "rabbitmq"
    vars.rabbitmq_check = "messages"
    assign where host.vars.rabbitmq_server.active
}

apply Service "rabbitmq-node" {
    check_command = "rabbitmq"
    vars.rabbitmq_check = "node"
    assign where host.vars.rabbitmq_server.active
}

apply Service "rabbitmq-ping" {
    check_command = "rabbitmq"
    vars.rabbitmq_check = "ping"
    assign where host.vars.rabbitmq_server.active
}

apply Service "rabbitmq-queues" {
    check_command = "rabbitmq"
    vars.rabbitmq_check = "queues"
    assign where host.vars.rabbitmq_server.active
}
```

### Host

```
object Host host use(host) {
  address = host
  ...
  
  vars.rabbitmq_server.active = "true"

  vars.rabbitmq_address = "https://" + address + ":15671"
  vars.rabbitmq_ca = "/etc/rabbitmq/legacy/ca.pem"
  vars.rabbitmq_cert = "/etc/rabbitmq/legacy/cert.pem"
  vars.rabbitmq_key = "/etc/rabbitmq/legacy/key.pem"
  vars.rabbitmq_password_file = "/etc/icinga2/secrets/rabbitmq"
  vars.rabbitmq_vhost = "production"
  vars.rabbitmq_node = "rabbit@" + address
  vars.rabbitmq_total_messages_warn_limit = 250000
  vars.rabbitmq_total_messages_critical_limit = 500000
}
```