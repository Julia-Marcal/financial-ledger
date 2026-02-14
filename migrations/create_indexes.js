// NOTE: MongoDB automatically creates a unique index on `_id` for every collection.
// Creating another _id index or specifying `unique: true` for the _id index will
// fail. Do not create explicit _id unique indexes.

// Index for fast balance calculation
db.transactions.createIndex({ accountId: 1 });

// Compound index for statement queries
db.transactions.createIndex({ accountId: 1, createdAt: -1 });

// The `_id` field is unique by default which provides idempotency guarantees for
// processed events. No explicit index creation for `_id` is necessary.
