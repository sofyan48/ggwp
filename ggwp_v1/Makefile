# Build And Development
test:
	@go test -v -cover ./...

depend:
	@go get github.com/ralali/event-api

devel:
	@air -c watcher.conf

build:
	@go build -o bgin src/main.go

run:
	@go run src/main.go

ggwp:
	@go run src/main.go -e production

worker:
	@go run src/main.go -e production

docker-stop:
	@docker-compose down

docker-image: 
	@docker build . -t bigevent_api

docker-run:
	@docker-compose up

.PHONY: test depend build  run stop docker docker-stop docker-image docker-run devel worker ggwp
