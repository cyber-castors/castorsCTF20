FROM ubuntu

RUN apt-get update && apt-get install socat -y
RUN mkdir data

COPY fmtme /data
COPY flag.txt /data

WORKDIR /data

EXPOSE 8080

# CMD ncat -k -e /data/fmtme -l 8080
CMD socat TCP-LISTEN:8080,reuseaddr,fork EXEC:./fmtme,pty,stderr,setsid,sigint,sane