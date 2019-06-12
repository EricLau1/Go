Ref: https://grpc.io/docs/quickstart/go/

go get -u google.golang.org/grpc

go get -u github.com/gin-gonic/gin

go get -u github.com/golang/protobuf/protoc-gen-go

export PATH=$PATH:$GOPATH/bin

protoc --proto_path=proto --go_out=plugins=grpc:proto proto/service.proto
protoc --proto_path=proto --go_out=plugins=grpc:proto proto/*.proto


