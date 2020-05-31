FROM ubuntu:19.10

RUN apt-get update
RUN apt-get install -y socat

RUN useradd -d /home/ctf/ -m -p ctf -s /bin/bash ctf

RUN echo "ctf:ctf" | chgpasswd

WORKDIR /home/ctf

COPY reverse_me .
COPY flag.txt .

RUN chown -R root:root /home/ctf

USER ctf

EXPOSE 8080

ENTRYPOINT socat TCP-LISTEN:8080,fork,reuseaddr exec:./reverse_me