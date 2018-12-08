run: docker-compose.yml
	docker-compose rm -f mongo redis
	docker-compose up --exit-code-from users

test-integration: docker-compose-test.yml
	docker-compose rm -f mongo redis
	docker-compose -f docker-compose-test.yml up --exit-code-from users

default: run