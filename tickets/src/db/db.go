package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//Client is a mongo client
var Client *mongo.Client

//ConnectDb connnect to the databse
func ConnectDb() {

	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Println("Error", err)

	}
	fmt.Println("Connected to MongoDB")
	// defer func() {
	// 	if err = Client.Disconnect(ctx); err != nil {
	// 		log.Println("Error: ", err)
	// 	}
	// }()

	t := Client.Ping(context.TODO(), readpref.Primary())
	if t == nil {
		fmt.Println("Connected To Client")

	}

}
