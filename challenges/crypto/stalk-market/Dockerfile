FROM python:3.7-alpine

RUN mkdir chal

COPY server.py /chal
COPY secret.py /chal
WORKDIR /chal

EXPOSE 8080

CMD python3 server.py