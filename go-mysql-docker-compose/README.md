MySQL Docker

docker build -t huskyh/docker-mysql:v1.0 db/

docker run -d -p 3306:3306 huskyh/docker-mysql:v1.0

docker exec -it [CONTAINER NAME] bash


App Docker

docker build -t huskyh/docker-golang:v1.0 app/

docker run -it --rm huskyh/docker-golang:v1.0