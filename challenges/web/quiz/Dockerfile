FROM golang:latest

RUN mkdir /quiz

COPY files /quiz/files
COPY main.go /quiz/main.go
COPY problems /quiz/problems

RUN mkdir -p /home/jeff

COPY flag.txt /home/jeff/flag.txt

RUN go get github.com/julienschmidt/httprouter

WORKDIR /quiz
EXPOSE 8080
CMD go run main.go