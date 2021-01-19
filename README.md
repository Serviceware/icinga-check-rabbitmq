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
