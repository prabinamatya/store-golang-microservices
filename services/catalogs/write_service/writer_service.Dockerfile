FROM golang:1.19.8

WORKDIR /app

ENV CONFIG=docker

COPY .. /app

RUN go get github.com/githubnemo/CompileDaemon
RUN go mod download


ENTRYPOINT CompileDaemon --build="go build -o main cmd/main.go" --command=./main