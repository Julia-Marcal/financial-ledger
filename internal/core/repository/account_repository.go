package mongodb

import (
	"context"
	"time"

	"financial-ledger/internal/core/model"
)

type AccountRepository interface {
	InsertAccount(ctx context.Context, acc model.Account) error
	FindAccountByID(ctx context.Context, id string) (model.Account, error)
	UpdateAccount(ctx context.Context, acc model.Account) error
	ListAccounts(ctx context.Context) ([]model.Account, error)
	GetBalance(ctx context.Context, accountId string) (int64, error)
	GetStatement(ctx context.Context, accountId string, from, to *time.Time) ([]model.Transaction, error)
}
