package service

import (
	"context"
	"time"

	"financial-ledger/internal/core/model"
	"financial-ledger/internal/infrastructure/mongodb"
)

func GetBalance(ctx context.Context, accountId string) (int64, error) {
	txs, err := mongodb.ListTransactions(ctx, accountId, nil, nil)
	if err != nil {
		return 0, err
	}
	var sum int64 = 0
	for _, t := range txs {
		if t.Type == "credit" {
			sum += t.Amount
		} else {
			sum -= t.Amount
		}
	}
	return sum, nil
}

func GetStatement(ctx context.Context, accountId string, from, to *time.Time) ([]model.Transaction, int64, error) {
	txs, err := mongodb.ListTransactions(ctx, accountId, from, to)
	if err != nil {
		return nil, 0, err
	}
	var balance int64 = 0
	for _, t := range txs {
		if t.Type == "credit" {
			balance += t.Amount
		} else {
			balance -= t.Amount
		}
	}
	return txs, balance, nil
}
