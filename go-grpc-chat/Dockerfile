FROM golang:alpine as build-env

ENV GO111MODULE=on

RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev

RUN mkdir /go-grpc-chat
RUN mkdir -p /go-grpc-chat/proto

WORKDIR /go-grpc-chat

COPY ./proto/service.pb.go /go-grpc-chat/proto
COPY ./main.go /go-grpc-chat

COPY go.mod .
COPY go.sum .

RUN go mod download 

RUN go build -o main .

CMD ./main