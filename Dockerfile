FROM golang:1.14-alpine

ADD . /go/src/github.com/VeloxMC/player-cache

RUN go install github.com/VeloxMC/player-cache
ENTRYPOINT /go/bin/player-cache

EXPOSE 8080