AWS SDK PARA GO

https://aws.amazon.com/pt/sdk-for-go/

Dep:

    go get  github.com/aws/aws-sdk-go

    go get  github.com/aws/aws-sdk-go/aws

    go get  github.com/aws/aws-sdk-go/service/cognitoidentity

    go get  github.com/aws/aws-sdk-go/service/dynamodb

    go get  github.com/aws/aws-sdk-go/service/s3

Configurar credênciais no arquivo:

    Linux or MacOS:

    ~/.aws/credentials

    Windows:

    c:\Users\username\.aws\credentials

Conteúdo do arquivo Credentials:

    [default]
    aws_access_key_id = YOUR_ACCESS_KEY
    aws_secret_access_key = YOUR_SECRET_KEY


Cognito:

    Pesquise por cognito dentro do console aws.

    Clique em gerenciar grupos de usuários.

    Clique em criar grupo de usuários.

    Crie um nome para o grupo.

    Clique em Revisar padrões.

    Na aba de atributos clique no icone de lápis para editar.

    Na seção 'Quais atributos padrão você deseja exigir?', selecione a opção 'name'.

    Clique em próxima etapa.

    Na página de políticas não alterar nada.

    Clique em próxima etapa.

    Em verificações e MFA não alterar nada.

    Clique em próxima etapa.

    Na seção 'Deseja personalizar suas mensagens de e-mail de verificação?', clique na opção 'LINK'.

    Clique em próxima etapa.

    Na página Tags, clique em próxima etapa.

    Na página Dispositivos, selecione Não, e clique em próxima etapa.

    Em Clientes do Aplicativo, adicione um cliente.

    Coloque um nome e clique em criar cliente.

    Clique em próxima etapa.

    Em Gatilhos clique em próxima etapa.

    Em revisar clique em criar grupo.

    Clique em Usuários e Grupos

    Crie um usuário, e informe o email e a senha de usuários da AWS.

    

