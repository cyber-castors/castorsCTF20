FROM ubuntu

RUN apt-get update && apt-get install socat -y
RUN mkdir data

COPY format /data
COPY flag.txt /data

WORKDIR /data

EXPOSE 8080

# CMD ncat -k -e /data/format -l 8080
#CMD socat TCP-LISTEN:8080,reuseaddr,fork EXEC:./format,pty,stderr,setsid,sigint,sane
CMD socat TCP-LISTEN:8080,reuseaddr,fork EXEC:./format,pty,stderr,setsid,sigint,sane,raw,echo=0