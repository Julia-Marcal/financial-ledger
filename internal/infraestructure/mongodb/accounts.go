package mongodb

import (
	"context"
	"financial-ledger/internal/core/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertAccount(ctx context.Context, acc model.Account) error {
	return WithClient(ctx, func(client *mongo.Client) error {
		coll := client.Database("ledger").Collection("accounts")

		doc := bson.M{
			"_id":       acc.ID,
			"createdAt": acc.Audit.CreatedAt,
			"audit": bson.M{
				"created": bson.M{"timestamp": acc.Audit.CreatedAt, "userId": acc.Audit.CreatedBy},
			},
		}

		if !acc.Audit.UpdatedAt.IsZero() || acc.Audit.UpdatedBy != "" {
			doc["audit"].(bson.M)["updated"] = bson.M{"timestamp": acc.Audit.UpdatedAt, "userId": acc.Audit.UpdatedBy}
		}

		_, err := coll.InsertOne(ctx, doc)
		return err
	})
}

func GetAccount(ctx context.Context, id string) (model.Account, error) {
	var acc model.Account
	err := WithClient(ctx, func(client *mongo.Client) error {
		coll := client.Database("ledger").Collection("accounts")
		return coll.FindOne(ctx, bson.M{"_id": id}).Decode(&acc)
	})
	return acc, err
}

func ListAccounts(ctx context.Context) ([]model.Account, error) {
	var accounts []model.Account
	err := WithClient(ctx, func(client *mongo.Client) error {
		coll := client.Database("ledger").Collection("accounts")
		findOpts := options.Find()
		findOpts.SetSort(bson.D{{Key: "createdAt", Value: -1}})

		cursor, err := coll.Find(ctx, bson.M{}, findOpts)
		if err != nil {
			return err
		}
		defer cursor.Close(ctx)

		if err := cursor.All(ctx, &accounts); err != nil {
			return err
		}
		return nil
	})
	return accounts, err
}
