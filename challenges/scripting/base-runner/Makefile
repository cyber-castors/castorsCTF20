.DEFAULT_GOAL := run

docker-run = docker run \
    --name base-runner -p 14430:8080

.PHONY: run
run: setup
    $(docker-run) 

.PHONY: setup
setup:
    @docker build . -t dot