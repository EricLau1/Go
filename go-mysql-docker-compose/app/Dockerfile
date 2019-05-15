FROM golang

WORKDIR /usr/src/app

RUN go get github.com/go-sql-driver/mysql

COPY . .

RUN go build -o main

EXPOSE 9000

CMD [ "./main" ]