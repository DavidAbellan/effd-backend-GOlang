package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoCn = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017/efdd_database")

/*connectDB Conecta a bd mongo*/
func ConnectDB() *mongo.Client {
	client, error := mongo.Connect(context.TODO(), clientOptions)
	if error != nil {
		log.Fatal(error.Error())
		return client
	}
	error = client.Ping(context.TODO(), nil)
	if error != nil {
		log.Fatal(error.Error())
		return client
	}
	log.Println("Connected!!")
	return client
}

/*checkConnection hace un ping a la bd*/
func CheckConnection() int {
	error := mongoCn.Ping(context.TODO(), nil)
	if error != nil {
		return 0
	}
	return 1
}
