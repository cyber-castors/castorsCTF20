.DEFAULT_GOAL := run

docker-run = docker run \
    --name arithmetics -p 14428:8080

.PHONY: run
run: setup
    $(docker-run) 

.PHONY: setup
setup:
    @docker build . -t dot