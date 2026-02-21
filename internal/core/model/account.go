package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID    string `json:"id" bson:"_id"`
	Audit Audit  `json:"audit" bson:"audit"`
}

func NewID() string {
	return primitive.NewObjectID().Hex()
}
