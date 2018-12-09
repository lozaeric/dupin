run: docker-compose.yml
	docker-compose rm -f mongo redis
	docker-compose up --abort-on-container-exit

clean: docker-compose.yml
	docker-compose rm -f mongo redis auth users messages

test-integration: docker-test-messages.yml docker-test-users.yml
	docker-compose rm -f mongo redis
	docker-compose -f docker-test-auth.yml up --exit-code-from auth
	docker-compose -f docker-test-messages.yml up --exit-code-from messages
	docker-compose -f docker-test-users.yml up --exit-code-from users

default: run