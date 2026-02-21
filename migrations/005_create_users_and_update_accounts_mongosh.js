// 005_create_users_and_update_accounts_mongosh.js
// Migration: Cria a coleção de usuários e adiciona campos de auditoria e users em accounts

// Cria a coleção de usuários se não existir
if (!db.getCollectionNames().includes("users")) {
  db.createCollection("users");
  db.users.createIndex({ email: 1 }, { unique: true });
}

// Adiciona campos em accounts (se não existirem)
db.accounts.updateMany(
  {},
  {
    $set: {
      createdAt: null,
      createdBy: null,
      updatedAt: null,
      updatedBy: null,
      users: []
    }
  }
);
