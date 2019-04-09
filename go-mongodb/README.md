Golang com MongoDB

Dependências via Glide:

    glide init 
    
    (Y)

    glide get gopkg.in/mgo.v2

Iniciar o cliente do mongodb e criar um banco que será usado neste exemplo.

Exportar duas variáveis de ambiente:

export MONGO_HOST=localhost && export MONGO_DB_NAME=mgo-demo

execute o comando => go fmt src/modules/profile/model/profile.go

Referências:

    https://gopkg.in/mgo.v2




