# create user
curl -XPOST localhost:8082/users -d '{"name":"eric","last_name":"loza","email":"lozaeric@gmail.com","password":"1234"}'
# create other user
curl -XPOST localhost:8082/users -d '{"name":"arthur","last_name":"pym","email":"lozaeric@pymtech.com","password":"1234"}'
# get token
curl -XPOST localhost:8081/token -d 'client_id=123123123&client_secret=111222333&username=bhbnpm743dd6118n2mig&password=1234&grant_type=client_credentials'
# send a message
curl -XPOST localhost:8080/messages -d '{"receiver":"bhbnsp743dd6118n2mjg","text":"hola mundo!"}' -H "x-auth:$TOKEN"
# search unseen messages