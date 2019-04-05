Golang Echo

Inicio

    Estrutura de diretórios:

    $ mkdir -p myapp/src/main

Configurar um GOPATH temporário a partir do diretório atual

    $ export GOPATH=`pwd`


Echo Framework 

    Site: 
    
    https://echo.labstack.com/

    Github:

    https://github.com/labstack/echo


    Instalar pacote:

    go get -u github.com/labstack/echo


Download Packages:

    go get github.com/labstack/echo/middlewares

    go get github.com/dgrijalva/jwt-go


Routes:

    GET     /

        *Hello World!*

    GET     /books/string?title=NewBook&author=Jane

        *Trocar o tipo da resposta para JSON: /books/json?title=NewBook&author=Jane*

    POST    /books

        *Envie no body da requisição um JSON com: { title: "The New Book", "author": "Jane Doe" }*

    POST    /contacts

        *Envie no body da requisição um JSON com: { "email": "email@email.com" }*

    GET     /jwt/login?username=admin&password=123456

        *return token JWT*

    GET     /jwt/main

        *Precisa do Header Authorization Bearer JWT*

    GET     /admin/main

        *Basic Auth Middleware. Entre com: admin, password: admin*

    GET     /login?username=admin&password=123456

    GET     /cookie/main
    
        *Precisa de autenticação por cookie. Utilize a penultima rota.*

Para Executar os arquivos estáticos:

    $ go install main

    $ cp bin/main static/

    $ ./static/main
