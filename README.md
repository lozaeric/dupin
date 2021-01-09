# Dupin

### Requirements
* Docker 1.13.0+
* Docker-compose

### Tech stack
* Golang 1.9.2
* MongoDB 3.6.9
* Redis 4.0.11

### Run integration tests
```bash
$ make test-integration
```

### Run
```bash
$ make run
```

### How to send and search a message
```bash
# creates user
curl -XPOST localhost:8082/users -d '{"name":"eric","last_name":"loza","email":"lz@dupin.com","password":"1234"}'
# creates another user
curl -XPOST localhost:8082/users -d '{"name":"arthur","last_name":"pym","email":"admin@dupin.com","password":"12345"}'
# gets an access token
curl -XPOST localhost:8081/token -d "client_id=123123123&client_secret=111222333&username=$USERID&password=$PASSWORD&grant_type=password"
# sends a message
curl -XPOST localhost:8080/messages -d '{"receiver_id":"$USERID","text":"hello world!"}' -H "x-auth:$TOKEN"
# searches messages
curl -XGET localhost:8080/search/messages -H "x-auth:$TOKEN"
```

### APIs
* messages
* users
* auth
* metric-collector
* groups (TODO)

### Endpoints
TODO
