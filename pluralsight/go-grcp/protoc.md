export PATH=$PATH:$GOPATH/bin

Go:

    protoc pb/messages.proto  --go_out=plugins=grpc:go/src


NodeJs:

    cd nodejs

    npm init -y

    npm i grpc @grpc/proto-loader  --save

Ref

    https://www.npmjs.com/package/grpc