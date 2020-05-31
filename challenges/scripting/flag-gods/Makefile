.DEFAULT_GOAL := run

docker-run = docker run \
    --name flag-gods -p 14431:8080

.PHONY: run
run: setup
    $(docker-run) 

.PHONY: setup
setup:
    @docker build . -t dot