package controllers

import "go.mongodb.org/mongo-driver/mongo"

// Goal -> Make a database connection
// Steps:
// 1- Define some constant variables in global scope
// 		1.1 Connection string
// 		1.2 DB name
// 		1.3 Collection name -> that's how mongodb works

// Most Important
// 2- We need to define collection variable and pass the reference of mongo db. for example
var collection *mongo.Collection

// Above will import the mongo driver

// 3- Now connect with mongo db
// 		This is going to be a method
// 		This will not be the main method but init(). For example:
// 		This method runs first time and only one at a time.
// 		Also called "initialization method"
func init() {
	// 1- First thing which we need to provide is "client" options.

}
