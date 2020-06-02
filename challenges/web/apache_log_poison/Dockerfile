FROM php:7.2-apache

COPY php /var/www/html
RUN mkdir -p /home/falg/flag/test/why/the/hassle/right
COPY flag.txt /home/falg/flag/test/why/the/hassle/right/flag.txt

RUN unlink /var/log/apache2/access.log
