package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID    string `json:"id" bson:"_id"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Audit Audit  `json:"audit" bson:"audit"`
}

// NewUserID deve ser implementado usando um gerador de UUID (ex: github.com/google/uuid)
func NewUserID() string {
	return primitive.NewObjectID().Hex()
}
