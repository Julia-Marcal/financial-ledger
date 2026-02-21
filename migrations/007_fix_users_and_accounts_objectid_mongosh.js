// 007_fix_users_and_accounts_objectid_mongosh.js
// Corrige o tipo dos campos de id para ObjectId em users e accounts

// Atualiza os ids dos usuários para ObjectId se necessário
const users = db.users.find({}).toArray();
users.forEach(user => {
  if (typeof user._id === 'string') {
    const newId = ObjectId();
    db.users.updateOne({ _id: user._id }, { $set: { _id: newId } });
    // Atualiza referências em accounts
    db.accounts.updateMany({ createdBy: user._id }, { $set: { createdBy: newId } });
    db.accounts.updateMany({ updatedBy: user._id }, { $set: { updatedBy: newId } });
    db.accounts.updateMany({ users: user._id }, { $set: { 'users.$': newId } });
  }
});

// Atualiza os ids das contas para ObjectId se necessário
const accounts = db.accounts.find({}).toArray();
accounts.forEach(acc => {
  if (typeof acc._id === 'string') {
    const newId = ObjectId();
    db.accounts.updateOne({ _id: acc._id }, { $set: { _id: newId } });
  }
});
