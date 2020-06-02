from golang:latest

RUN apt-get update -qy 

RUN mkdir /opt/xss
COPY controllers /opt/xss/controllers
COPY models /opt/xss/models
COPY templates /opt/xss/templates
COPY flag.txt /opt/xss
COPY main.go /opt/xss

WORKDIR /opt/xss

EXPOSE 8080
CMD go run main.go