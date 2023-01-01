package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// "go.mongodb.org/mongo-driver/bson/primitive"
// This package has to be imported while working with mongo db
// bson is similar to json
// BSON is a binary serialization format used to store documents and make remote procedure calls in MongoDB. The BSON specification is located at bsonspec.org

type Netflix struct {
	ID      primitive.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}

// primitive.ObjectId
// about to learn
