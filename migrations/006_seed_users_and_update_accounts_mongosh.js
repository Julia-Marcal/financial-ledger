function uuidv4() {
  return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
    var r = Math.random() * 16 | 0, v = c == 'x' ? r : (r & 0x3 | 0x8);
    return v.toString(16);
  });
}

const userId = uuidv4();
db.users.insertOne({
  _id: userId,
  name: "System2",
  email: "system2@ledger.local",
  createdAt: new Date(),
  updatedAt: new Date()
});