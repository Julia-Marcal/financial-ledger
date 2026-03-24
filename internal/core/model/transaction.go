package model

import (
	"time"
)

type Transaction struct {
	AccountID      string    `json:"account_id" bson:"account_id"`
	Type           string    `json:"type" bson:"type"` // credit || debit
	Amount         int64     `json:"amount" bson:"amount"`
	CreatedAt      time.Time `json:"createdAt" bson:"createdAt"`
	IdempotencyKey string    `json:"idempotency_key" bson:"idempotency_key"`
}
