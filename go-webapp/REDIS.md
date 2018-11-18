Instalar Redis no Linux:

    sudo apt-get install redis

Comandos Básicos

Iniciar o client:

$ redis-cli

Adicionando uma Lista de comentarios:

$ lpush [chave] [valor]

 *lpush e o comando que ira adicionar um valor a lista referente a chave*

Exemplo:

$ lpush comentarios "Hello world"

Como ler a lista de comentários:

$ lrange [chave] [inicio] [termino]

*os parametros inicio e termino indicam a quantidade de valores que serão mostrados da lista reference a chave*

Exemplo:

$ lrange users 0 10

*neste exemplo serão mostrados apenas os 10 primeiros usuários*

Removendo elemento da lista:

$ lrem [chave] [contador] [valor]

*remove a quantidade de elementos que for indicada no contador de acordo com o valor fornecido referente a chave*


Removento a lista completa

$ del [chave]

Visualizar todas as Keys

$ keys *

Visualizar Keys que tenham uma palavra específica

    $ keys *palavra*

Exemplo:

    $ keys *user:*

Comandos específicos deste projeto

visualizar chave simples

$ get [chave]

    exemplo:

        $ get "user:jane"

visualizar campos de um hashmap:

$ hget [chave] [campo]

    exemplo:

        $ hget "user:1" "username"

visualizar todos os campos de hashmap:

$ hgetall [chave]

    exemplo:

        $ hgetall "user:1"

Destruindo todas as chaves:

    $ flushdb

Parando o servidor:

$ sudo service redis stop

Iniciando o servidor

$ sudo service redis start



