package main

import (
	"github.com/lozaeric/dupin/app"
)

func main() {
	app.Run()
}

/*
func testing() {
	connectionString := "mongodb://root:example@mongo:27017/"

	client, err := mongo.NewClient(connectionString)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(nil, nil)
	if err != nil {
		log.Fatal(err)
	}
}
*/
