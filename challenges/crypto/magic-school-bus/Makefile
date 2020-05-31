.DEFAULT_GOAL := run

docker-run = docker run \
    --name magic-school-bus -p 14435:8080

.PHONY: run
run: setup
    $(docker-run) 

.PHONY: setup
setup:
    @docker build . -t dot