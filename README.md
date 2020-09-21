# Dupin

### Technologies
* Golang 1.9.2
* MongoDB 3.6.9
* Redis 4.0.11

### Run
```bash
$ make run
```

### Integration test
```bash
$ make test-integration
```

## Manual integration test
```bash
# create user
curl -XPOST localhost:8082/users -d '{"name":"eric","last_name":"loza","email":"lz@pymtech.com","password":"1234"}'
# create other user
curl -XPOST localhost:8082/users -d '{"name":"arthur","last_name":"pym","email":"admin@pymtech.com","password":"12345"}'
# get token
curl -XPOST localhost:8081/token -d "client_id=123123123&client_secret=111222333&username=$USERID&password=$PASS&grant_type=password"
# send a message
curl -XPOST localhost:8080/messages -d '{"receiver":"$USERID","text":"hola mundo!"}' -H "x-auth:$TOKEN"
# search messages
curl -XGET localhost:8080/search/messages -H "x-auth:$TOKEN"
```

### APIs
* messages-api
* users-api
* auth-api
* storage-api (TODO)

### Endpoints
TODO
