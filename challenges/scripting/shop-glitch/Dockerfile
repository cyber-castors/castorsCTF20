FROM golang:latest

RUN mkdir /shop

COPY main.go /shop
COPY flag.txt /shop

WORKDIR /shop

EXPOSE 8080
CMD go run main.go