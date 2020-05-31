.DEFAULT_GOAL := run

docker-run = docker run \
    --name shop-glitch -p 14432:8080

.PHONY: run
run: setup
    $(docker-run) 

.PHONY: setup
setup:
    @docker build . -t dot