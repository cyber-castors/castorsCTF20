.DEFAULT_GOAL := run

docker-run = docker run \
    --name one-trick-pony -p 14422:8080

.PHONY: run
run: setup
    $(docker-run) 

.PHONY: setup
setup:
    @docker build . -t dot