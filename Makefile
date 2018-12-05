run: docker-compose.yml
	docker-compose rm -f mongo redis
	docker-compose up --abort-on-container-exit --exit-code-from web

test-integration: docker-compose-test.yml
	docker-compose rm -f mongo redis
	docker-compose -f docker-compose-test.yml up --abort-on-container-exit --exit-code-from web

default: run