clean: docker-compose.yml
	make clean-services
	docker-compose rm -f auth users messages

clean-services:
	docker-compose rm -f mongo redis

run: docker-compose.yml
	docker-compose up --abort-on-container-exit

test-integration: docker-test-auth.yml docker-test-messages.yml docker-test-users.yml
	make clean-services
	make test-auth
	make test-messages
	make test-users

test-auth: docker-test-auth.yml
	docker-compose -f docker-test-auth.yml up --exit-code-from auth

test-messages: docker-test-messages.yml
	docker-compose -f docker-test-messages.yml up --exit-code-from messages

test-users: docker-test-users.yml
	docker-compose -f docker-test-users.yml up --exit-code-from users

default: run