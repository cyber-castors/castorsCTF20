.DEFAULT_GOAL := run

docker-run = docker run \
    --name bagel-bytes -p 14420:8080

.PHONY: run
run: setup
    $(docker-run) 

.PHONY: setup
setup:
    @docker build . -t dot