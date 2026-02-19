package mongodb

import (
	"context"
	"time"

	"financial-ledger/internal/core/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

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
