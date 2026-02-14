if (db.getCollectionInfos({ name: "processed_events" }).length === 0) {
    db.createCollection("processed_events", {
        validator: {
            $jsonSchema: {
                bsonType: "object",
                required: ["_id", "processedAt"],
                properties: {
                    _id: {
                        bsonType: "string",
                        description: "UUID do evento"
                    },
                    processedAt: {
                        bsonType: "date"
                    }
                }
            }
        }
    });
} else {
    print("Collection 'processed_events' already exists - skipping");
}
