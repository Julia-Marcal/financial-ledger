if (db.getCollectionInfos({ name: "accounts" }).length === 0) {
    db.createCollection("accounts", {
        validator: {
            $jsonSchema: {
                bsonType: "object",
                required: ["_id", "createdAt", "audit"],
                properties: {
                    _id: {
                        bsonType: "string",
                        description: "UUID da conta - obrigatório"
                    },
                    createdAt: {
                        bsonType: "date",
                        description: "Data de criação - obrigatório"
                    },
                    audit: {
                        bsonType: "object",
                        required: ["created"],
                        properties: {
                            created: {
                                bsonType: "object",
                                required: ["timestamp"],
                                properties: {
                                    timestamp: { bsonType: "date", description: "creation timestamp - required" },
                                    userId: { bsonType: "string", description: "user who created the document" }
                                }
                            },
                            updated: {
                                bsonType: "object",
                                properties: {
                                    timestamp: { bsonType: "date", description: "last update timestamp" },
                                    userId: { bsonType: "string", description: "user who last updated" }
                                }
                            }
                        }
                    }
                }
            }
        }
    });
} else {
    print("Collection 'accounts' already exists - skipping");
}
