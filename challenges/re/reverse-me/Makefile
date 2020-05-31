.DEFAULT_GOAL := run

docker-run = docker run \
    --name reverse-me -p 14427:8080

.PHONY: run
run: setup
    $(docker-run) 

.PHONY: setup
setup:
    @docker build . -t dot