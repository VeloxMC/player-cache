FROM golang:1.14-alpine

ADD . /go/src/github.com/VeloxMC/player-cache

RUN go run github.com/VeloxMC/player-cache/main.go

EXPOSE 8080