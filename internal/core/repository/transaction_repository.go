package mongodb

import (
	"context"
	"time"

	"financial-ledger/internal/core/model"
)

type TransactionRepository interface {
	InsertTransaction(ctx context.Context, tx model.Transaction) error
	ListTransactionsWithFilter(ctx context.Context, accountId string, from, to *time.Time) ([]model.Transaction, error)
	ListTransactions(ctx context.Context, accountId string, from, to *time.Time) ([]model.Transaction, error)
}
