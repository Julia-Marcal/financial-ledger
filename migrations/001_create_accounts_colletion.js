if (db.getCollectionInfos({ name: "accounts" }).length === 0) {
    db.createCollection("accounts", {
        validator: {
            $jsonSchema: {
                bsonType: "object",
                required: ["_id", "createdAt"],
                properties: {
                    _id: {
                        bsonType: "string",
                        description: "UUID da conta - obrigatório"
                    },
                    createdAt: {
                        bsonType: "date",
                        description: "Data de criação - obrigatório"
                    }
                }
            }
        }
    });
} else {
    print("Collection 'accounts' already exists - skipping");
}
