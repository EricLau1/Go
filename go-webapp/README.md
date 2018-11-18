Download packages

    go get github.com/gorilla/mux
 
    go get github.com/gorilla/sessions

    go get github.com/go-redis/redis

    go get golang.org/x/crypto/bcrypt 


Estrutura do projeto

[models] 
    arquivos de conexão com o banco de dados

[sessions]
    arquivo de sessão 

[static]
    arquivos estáticos -> css, js

[templates]
    páginas html

[utils]
    arquivos para carregar templates (os arquivos html - views)

[routes]
    arquivo para definir o caminho das requisições pela URL

[middleware]
    arquivo que define quais páginas são necessárias ter uma sessão aberta (usuário logado)

'main.go'
    executa o app

