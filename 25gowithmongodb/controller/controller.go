package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

// Goal -> Make a database connection
// Steps:
// 1- Define some constant variables in global scope
// 		1.1 Connection string
// 		1.2 DB name
// 		1.3 Collection name -> that's how mongodb works

const connectionString string = ""
const dbName string = "cs-go"
const collectionName string = "golangMongoDBCollection"

// Most Important
// 2- We need to define collection variable and pass the reference of mongo db. for example
var collection *mongo.Collection

// Above will import the mongo driver

// 3- Now connect with mongo db
// 		This is going to be a method
// 		This will not be the main method but init(). For example:
// 		This init() is a specialized method in golang which runs only at the very first time this entire application is going to execute.
// 		This method runs only first time and only one at a time.
// 		Also called "initialization method"
func init() {
	// 1- First thing which we need to provide is "client" options.
	clientOption := options.Client().ApplyURI(connectionString) // this client connection doesn't fire up any connection request

	// Connect to mongo db
	client, err := mongo.Connect(context.TODO(), clientOption) // We are using this TODO type (struct) because we are not sure which context we want to use.
	// In most of the cases, when we want to connect to mongo db and perform crud operations, we use Backgroud type
	// This client is now completely capable of connecting with the database.
	// Find more details below about context

	// context
	// 		context -> is a package in golang
	// 		In most of the cases, databases are outside of the web server (or we can say in another machine)
	// 		When this web server (or API) communicates to database, it requires a context in which both (API/Web server and database) can work together
	// 		In this case, we need to provide the context,For example,
	// 			- for how long we can make a request
	// 			- does our connection still established, if it is, then there should be a context in which we can work on with.
	// 			- is there any issue with the connection (failed or something else)
	// 		So, these kind of things should be done in any context
	// 	There are different types (struct) of context
	// 		- Background
	// 		- TODO -> this type (of context) is used when we are not clear that when to use which context
	// 		- WithValue

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDb connection successful!")
	// ---------------------------- We have done database connection --------------------------------------------

	// ---------------------------- Now, we will need to go inside the database and it's collection -------------
	// To do this, we need to provide the reference of collection so we don't need to define values again and again.
	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection instance is ready!")

}

// In this init() function, we've connected to database and went inside the db and collection.
// Now here is the action plan
// 1- Insert the data in mongodb
// 		1.1 Define two methods
// 			1.1.1 Mongodb helper method
// 					This method will insert the data in db
// 			1.1.2 Method to receive the data from request body
// 					This will receive the data from request body
// 					Perform all checks
// 					Call mongodb helper method to insert the data

// Create mongodb helper method
// Insert One Record

func insertOneRecord(movie model.Netflix) {

	inserted, err := collection.InsertOne(context.Background(), movie)

	checkNilErr(err)

	fmt.Println("Inserted one record in db ", inserted.InsertedID)

}

func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
