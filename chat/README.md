# CHAT CLIENT

## Getting Started
### Environment Setup
Copy .env.example to .env
```
KAFKA_PORT=9092
KAFKA_HOST=localhost
KAFKA_USERNAME=
KAFKA_PASSWORD=
```

### Started Window Chat
```
go run main.go
```
setup room and chat id

### Started Sending Window Chat
```
go run main.go -w send
```

### Example
![chat](https://github.com/sofyan48/ggwp/blob/master/client/docs/example.png)
