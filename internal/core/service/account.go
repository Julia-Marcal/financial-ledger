package service

import (
	"context"
	"time"

	"financial-ledger/internal/core/model"
	"financial-ledger/internal/infrastructure/mongodb"
)

func CreateAccount(ctx context.Context, acc model.Account) (model.Account, error) {
	if acc.ID == "" {
		acc.ID = model.NewID()
	}
	if acc.CreatedAt.IsZero() {
		acc.CreatedAt = time.Now().UTC()
	}

	if acc.Audit.Created.Timestamp.IsZero() {
		acc.Audit.Created.Timestamp = acc.CreatedAt
	}
	if acc.Audit.Updated.Timestamp.IsZero() {
		acc.Audit.Updated.Timestamp = acc.Audit.Created.Timestamp
	}

	err := mongodb.InsertAccount(ctx, acc)
	return acc, err
}

func GetAccount(ctx context.Context, id string) (model.Account, error) {
	return mongodb.GetAccount(ctx, id)
}

func ListAccounts(ctx context.Context) ([]model.Account, error) {
	return mongodb.ListAccounts(ctx)
}
