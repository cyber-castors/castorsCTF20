FROM ubuntu

RUN apt-get update && apt-get install ncat -y && adduser leet --uid 1337 --gecos "" --disabled-password --disabled-login
RUN mkdir data
COPY babybof /data
COPY flag.txt /data
COPY shell_flag.txt /data

USER leet

WORKDIR /data

EXPOSE 8080

CMD ncat -k -e /data/babybof -l 8080

