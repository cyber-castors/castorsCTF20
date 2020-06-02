FROM golang:latest

RUN apt-get update -qy && apt-get upgrade -qy
RUN go get github.com/julienschmidt/httprouter
RUN go get github.com/satori/go.uuid

RUN mkdir /opt/shortcuts
COPY shortcuts /opt/shortcuts/shortcuts
COPY controllers /opt/shortcuts/controllers
COPY main.go /opt/shortcuts/main.go
COPY models /opt/shortcuts/models
COPY templates /opt/shortcuts/templates

RUN mkdir -p /home/tom/
COPY flag.txt /home/tom/flag.txt

WORKDIR /opt/shortcuts/

#CMD [ "/bin/bash" ]
EXPOSE 8080
CMD [ "go","run","main.go" ]

