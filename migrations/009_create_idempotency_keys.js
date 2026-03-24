db.createCollection("idempotency_keys", {
    validator: {
        $jsonSchema: {
            bsonType: "object",
            required: ["_id", "status", "createdAt"],
            properties: {
                _id: {
                    bsonType: "string",
                    description: "Idempotency key — must be unique"
                },
                status: {
                    bsonType: "string",
                    enum: ["processed"],
                    description: "Processing status"
                },
                createdAt: {
                    bsonType: "date",
                    description: "Timestamp of first processing"
                }
            }
        }
    }
});
