Instalação do Protobuf

    *Tem que ter instalado o C++*

    sudo apt  install protobuf-compiler

Configurando o protobuf para Golang:

    go get -u -v github.com/golang/protobuf/protoc-gen-go

    export PATH=$PATH:$GOPATH/bin

    go get -u -v github.com/golang/protobuf/proto

Comando:

    protoc --go_out=. *.proto


Run:

   go run main.go *.pb.go


