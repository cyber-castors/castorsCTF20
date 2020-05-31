FROM ubuntu

RUN apt-get update && apt-get install socat -y
RUN mkdir data

COPY abcbof /data
COPY flag.txt /data

WORKDIR /data

EXPOSE 8080

# CMD ncat -k -e /data/abcbof -l 8080
CMD socat TCP-LISTEN:8080,reuseaddr,fork EXEC:./abcbof,pty,stderr,setsid,sigint,sane