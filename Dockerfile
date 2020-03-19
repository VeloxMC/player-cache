FROM golang:1.14-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o playercache .

EXPOSE 8080

CMD ["./playercache"]