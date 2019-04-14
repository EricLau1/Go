CURL:

    https://curl.haxx.se/docs/manpage.html


GET:

    curl -X GET http://localhost:8080

    Headers:

    curl -I -X GET http://localhost:8080

    All Headers:
    
    curl -v http://localhost:8080


POST:

    curl -d "email=email@email.com&password=123456" -X POST http://localhost:8080

    Headers:

    curl -i -d "email=email@email.com&password=123456" -X POST http://localhost:8080

    AUTH:

    curl -d "email=admin&password=admin" -X POST http://localhost:8080/admin


PUT:

    curl -d "email=email@email.com&password=123456" -X PUT http://localhost:8080?id=35254

    headers:

    curl -i -d "email=email@email.com&password=123456" -X PUT http://localhost:8080?id=35254

DELETE

    curl -X DELETE http://localhost:8080?id=123456

    headers:

    curl -i -X DELETE http://localhost:8080?id=123456


JSON:

    curl --header "Content-Type: application/json" \
         --request POST \
         --data "{\"email\":\"admin\",\"password\":\"123456\"}" \
         http://localhost:8080/json

