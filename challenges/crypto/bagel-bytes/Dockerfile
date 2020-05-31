FROM python:3.7-alpine

COPY requirements.txt /
RUN apk add gcc g++ make libffi-dev openssl-dev
RUN pip3 install -r requirements.txt
RUN mkdir chal

COPY server.py /chal
COPY secret.py /chal
WORKDIR /chal

EXPOSE 8080

CMD python3 server.py