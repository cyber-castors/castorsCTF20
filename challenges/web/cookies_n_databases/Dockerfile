FROM golang:latest
RUN apt-get update -qy && apt-get upgrade -qy

RUN mkdir /opt/cookies
COPY controllers /opt/cookies/controllers
COPY models /opt/cookies/models
COPY main.go /opt/cookies

RUN go get github.com/julienschmidt/httprouter
RUN go get github.com/go-sql-driver/mysql
RUN go get golang.org/x/crypto/bcrypt

WORKDIR /opt/cookies

EXPOSE 8080
CMD [ "go", "run", "main.go" ]
#CMD "/bin/bash"