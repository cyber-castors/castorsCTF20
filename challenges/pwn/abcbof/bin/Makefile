.DEFAULT_GOAL := run

docker-run = docker run \
	--name abcbof -p 14424:8080 \
	--rm \
	--volume $(shell pwd):/data/:ro

.PHONY: run
run: setup
	$(docker-run) 

.PHONY: setup
setup:
	@docker build . -t dot 
