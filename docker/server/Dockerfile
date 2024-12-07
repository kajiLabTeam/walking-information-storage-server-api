FROM golang:1.23.3

WORKDIR /server

COPY /server /server
COPY /.env /server/.env

CMD ["go", "run", "main.go"]
