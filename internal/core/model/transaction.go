package model

import (
	"time"
)

type Transaction struct {
	ID        string    `json:"id" bson:"_id"`
	AccountID string    `json:"accountId" bson:"accountId"`
	Type      string    `json:"type" bson:"type"` // credit || debit
	Amount    int64     `json:"amount" bson:"amount"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}
