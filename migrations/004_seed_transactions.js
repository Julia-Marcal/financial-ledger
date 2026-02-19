// Seed sample accounts (if missing) and transactions for development/testing.
// Safe to run multiple times: accounts are upserted and transactions are inserted
// only if the `transactions` collection is empty.

if (db.getCollectionInfos({ name: "transactions" }).length === 0) {
    throw new Error("Collection 'transactions' does not exist. Run migrations to create collections first.");
}

if (db.transactions.countDocuments({}) > 0) {
    print("Transactions collection is not empty - skipping seed")
} else {
    // Ensure some accounts exist (upsert so we don't duplicate)
    const sampleAccounts = [
        { _id: "acct-1111-aaaa", createdAt: new Date("2024-01-01T00:00:00Z") },
        { _id: "acct-2222-bbbb", createdAt: new Date("2024-02-01T00:00:00Z") },
        { _id: "acct-3333-cccc", createdAt: new Date("2024-03-01T00:00:00Z") }
    ];

    sampleAccounts.forEach(a => {
        db.accounts.updateOne({ _id: a._id }, { $setOnInsert: a }, { upsert: true });
    });

    // Create sample transactions.
    // `amount` is stored as cents (bsonType: long in the migration schema).
    const txs = [
        { _id: "tx-0001", accountId: "acct-1111-aaaa", type: "credit", amount: NumberLong("150000"), createdAt: new Date("2024-01-05T10:00:00Z"), description: "Initial deposit" },
        { _id: "tx-0002", accountId: "acct-1111-aaaa", type: "debit", amount: NumberLong("2500"), createdAt: new Date("2024-01-10T12:30:00Z"), description: "Coffee shop" },
        { _id: "tx-0003", accountId: "acct-1111-aaaa", type: "debit", amount: NumberLong("5000"), createdAt: new Date("2024-01-11T09:15:00Z"), description: "Groceries" },
        { _id: "tx-0004", accountId: "acct-2222-bbbb", type: "credit", amount: NumberLong("200000"), createdAt: new Date("2024-02-03T08:00:00Z"), description: "Salary" },
        { _id: "tx-0005", accountId: "acct-2222-bbbb", type: "debit", amount: NumberLong("12000"), createdAt: new Date("2024-02-04T14:00:00Z"), description: "Electric bill" },
        { _id: "tx-0006", accountId: "acct-2222-bbbb", type: "debit", amount: NumberLong("3000"), createdAt: new Date("2024-02-07T19:45:00Z"), description: "Restaurant" },
        { _id: "tx-0007", accountId: "acct-3333-cccc", type: "credit", amount: NumberLong("50000"), createdAt: new Date("2024-03-02T11:00:00Z"), description: "Gift" },
        { _id: "tx-0008", accountId: "acct-3333-cccc", type: "debit", amount: NumberLong("7500"), createdAt: new Date("2024-03-10T16:20:00Z"), description: "Books" },
        { _id: "tx-0009", accountId: "acct-1111-aaaa", type: "credit", amount: NumberLong("5000"), createdAt: new Date("2024-01-20T09:00:00Z"), description: "Refund" },
        { _id: "tx-0010", accountId: "acct-3333-cccc", type: "debit", amount: NumberLong("2000"), createdAt: new Date("2024-03-15T13:00:00Z"), description: "Taxi" }
    ];

    const result = db.transactions.insertMany(txs);
    // mongosh may return insertedIds instead of insertedCount; normalize
    let insertedCount = 0;
    if (result && result.insertedIds) {
        insertedCount = Object.keys(result.insertedIds).length;
    } else if (result && typeof result.insertedCount === 'number') {
        insertedCount = result.insertedCount;
    }
    print(`Inserted ${insertedCount} transactions`);
}
