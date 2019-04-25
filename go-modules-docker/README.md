Env:

    export GO111MODULE=on

Run:

    go mod init

    go build

Docker commands:

    Criar imagem:

    sudo docker build -t goapp .

    Rodar imagem:

    sudo docker run -it -p 9000:9001 goapp

    Rodar imagem em background:

    sudo docker run -d -p 9000:9001 goapp

    Visualizar imagens:

    sudo docker ps

    Parar imagem:

    sudo docker kill [DOCKER ID]