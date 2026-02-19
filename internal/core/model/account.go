package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID        string    `json:"id" bson:"_id"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	Audit     Audit     `json:"audit" bson:"audit"`
}

func NewID() string {
	return primitive.NewObjectID().Hex()
}
