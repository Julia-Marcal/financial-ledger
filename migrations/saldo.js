db.transactions.aggregate([
    {
        $match: { accountId: "ACCOUNT_UUID" }
    },
    {
        $group: {
            _id: null,
            balance: {
                $sum: {
                    $cond: [
                        { $eq: ["$type", "credit"] },
                        "$amount",
                        { $multiply: ["$amount", -1] }
                    ]
                }
            }
        }
    }
]);
