package mongodb

import (
	"context"
	"time"

	"financial-ledger/internal/core/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertTransaction(ctx context.Context, tx model.Transaction) error {
	return WithClient(ctx, func(client *mongo.Client) error {
		coll := client.Database("ledger").Collection("transactions")
		doc := bson.M{
			"_id":       tx.ID,
			"accountId": tx.AccountID,
			"type":      tx.Type,
			"amount":    tx.Amount,
			"createdAt": tx.CreatedAt,
		}
		_, err := coll.InsertOne(ctx, doc)
		return err
	})
}

func ListTransactionsWithFilter(ctx context.Context, accountId string, from, to *time.Time) ([]model.Transaction, error) {
	filter := bson.M{}
	if accountId != "" {
		filter["accountId"] = accountId
	}
	if from != nil || to != nil {
		rng := bson.M{}
		if from != nil {
			rng["$gte"] = *from
		}
		if to != nil {
			rng["$lte"] = *to
		}
		filter["createdAt"] = rng
	}
	var txs []model.Transaction
	err := WithClient(ctx, func(client *mongo.Client) error {
		coll := client.Database("ledger").Collection("transactions")
		cursor, err := coll.Find(ctx, filter)
		if err != nil {
			return err
		}
		defer cursor.Close(ctx)
		return cursor.All(ctx, &txs)
	})
	return txs, err
}
func ListTransactions(ctx context.Context, accountId string, from, to *time.Time) ([]model.Transaction, error) {
	var txs []model.Transaction
	err := WithClient(ctx, func(client *mongo.Client) error {
		coll := client.Database("ledger").Collection("transactions")
		filter := bson.M{"accountId": accountId}
		if from != nil || to != nil {
			rng := bson.M{}
			if from != nil {
				rng["$gte"] = *from
			}
			if to != nil {
				rng["$lte"] = *to
			}
			filter["createdAt"] = rng
		}

		cursor, err := coll.Find(ctx, filter)
		if err != nil {
			return err
		}
		defer cursor.Close(ctx)

		return cursor.All(ctx, &txs)
	})
	return txs, err
}
