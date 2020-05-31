.DEFAULT_GOAL := run

docker-run = docker run \
    --name stalk-market -p 14423:8080

.PHONY: run
run: setup
    $(docker-run) 

.PHONY: setup
setup:
    @docker build . -t dot