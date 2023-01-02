package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// "go.mongodb.org/mongo-driver/bson/primitive"
// This package has to be imported while working with mongo db
// bson is similar to json
// bson
// Whenever we use mongodb, it gives us the ID which mongo db generates automatically.
// These are not just IDs, these are of "_id" type. They are kind of a bson structure similar to json with some extended fields added to it.
// BSON is a binary serialization format used to store documents and make remote procedure calls in MongoDB. The BSON specification is located at bsonspec.org

type Netflix struct {
	ID      primitive.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}

// primitive.ObjectId
// This ObjectId is exactly the underscore id (_id) that is given to us
// This line ("go.mongodb.org/mongo-driver/bson/primitive") is responsible for primitive.ObjectId
