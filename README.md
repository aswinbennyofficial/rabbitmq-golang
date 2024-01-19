# rabbitmq-golang
Repo to experiment with RabbitMQ and golang

### Dependency
```bash
go get github.com/rabbitmq/amqp091-go
```

### Setup RabbitMQ server

```bash
docker run --rm -it -p 15672:15672 -p 5672:5672 rabbitmq:3-management
```

- 15672 : web management
- 5672 : broker
- Username and Password : `guest`

### Usage

- Run Producer : 
```bash
go run ./send/
```

- Run Consumer :
```bash
go run ./recieve/
```
