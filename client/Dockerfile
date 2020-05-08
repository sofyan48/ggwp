## Builder
FROM golang:latest

RUN mkdir -p /app

WORKDIR /app
COPY    . .
RUN     go build -o main main.go 

EXPOSE 3000