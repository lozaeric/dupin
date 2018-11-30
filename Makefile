run: docker-compose.yml
	docker-compose up

test:
	go test -v apitest/*

default: test