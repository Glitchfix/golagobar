use golagobar
db.createCollection("users")
db.createCollection("auth")
db.createCollection("rides")
db.rides.insertOne({
    "_id" : 0,
    "rideID" : "rideID",
    "sequence" : 1
})
db.users.createIndex({"location": "2dsphere"})
