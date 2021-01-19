FROM golang:1.14

WORKDIR /go/src/app
COPY main.go .
CMD go run main.go

