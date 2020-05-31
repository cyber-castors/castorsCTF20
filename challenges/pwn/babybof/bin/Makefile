.DEFAULT_GOAL := run

docker-run = docker run \
	--name babybof -p 14425:8080 \
	-d \
	--volume $(shell pwd):/data/:ro

.PHONY: run
run: setup
	$(docker-run) 

.PHONY: setup
setup:
	@docker build . -t dot
