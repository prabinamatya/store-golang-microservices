FROM golang:1.20.11

WORKDIR /app

ENV CONFIG=docker

COPY .. /app

RUN go get github.com/githubnemo/CompileDaemon
RUN go mod download


ENTRYPOINT CompileDaemon --build="go build -o main cmd/main.go" --command=./main