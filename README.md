# Icinga check for RabbitMQ

Icinga check for RabbitMQ

## Build

```
go build -o check_rabbitmq ./cmd/CheckRabbitMQ.go
```

## Test

Start a docker container for RabbitMQ:
```
docker run -p 5672:5672 -p 5671:5671 -p 15672:15672 -p 15671:15671 rabbitmq:3.8.10-rc.6-management
```

Run tests:
```
go test -v ./test
```
