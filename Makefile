run: docker-compose.yml
	docker-compose up

test-integration:
	docker-compose -f docker-compose-test.yml up --abort-on-container-exit

default: run