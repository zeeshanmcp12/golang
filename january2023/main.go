package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	fmt.Println("Working with Mongo db")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(""))
	checkNilErr(err)

	err = client.Ping(ctx, readpref.Primary())

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
