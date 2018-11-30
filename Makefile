run: docker-compose.yml
	docker-compose up

test:
	go test -v apitest/*

test-integration:
	docker-compose -f docker-compose-test.yml up

default: run