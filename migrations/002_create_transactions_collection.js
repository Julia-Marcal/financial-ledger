if (db.getCollectionInfos({ name: "transactions" }).length === 0) {
    db.createCollection("transactions", {
        validator: {
            $jsonSchema: {
                bsonType: "object",
                required: ["_id", "accountId", "type", "amount", "createdAt"],
                properties: {
                    _id: {
                        bsonType: "string",
                        description: "UUID da transação"
                    },
                    accountId: {
                        bsonType: "string",
                        description: "UUID da conta"
                    },
                    type: {
                        enum: ["credit", "debit"],
                        description: "Tipo da transação"
                    },
                    amount: {
                        bsonType: "long",
                        minimum: 1,
                        description: "Valor em centavos"
                    },
                    createdAt: {
                        bsonType: "date"
                    }
                }
            }
        }
    });
} else {
    print("Collection 'transactions' already exists - skipping");
}
