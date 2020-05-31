FROM ubuntu

RUN apt-get update && apt-get install socat -y && adduser leet --uid 1337 --gecos "" --disabled-password --disabled-login
RUN apt-get install libc6-i386 -qy
RUN mkdir data
COPY winners /data
COPY flag.txt /data

USER leet

WORKDIR /data

EXPOSE 8080

# CMD ncat -k -e ./winners -l 8080
CMD socat TCP-LISTEN:8080,reuseaddr,fork EXEC:./winners,pty,stderr,setsid,sigint,sane

