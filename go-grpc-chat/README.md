# Build and Run

    export PATH=$PATH:$GOPATH/bin

    export GO111MODULE=on

    go get -u -v google.golang.org/grpc/grpclog

    protoc --proto_path=proto --go_out=plugins=grpc:proto proto/*.proto

    go mod init

    go build -o main 

    ./main

## Docker

    docker build -t go-grpc-chat:v1 .

    docker run -it --rm -p 8080:8080 go-grpc-chat:v1


## Client Server

    go run client/main.go -N Username


