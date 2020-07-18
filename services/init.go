package services

import (
	"github.com/Glitchfix/golagobar/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	userCollection  *mongo.Collection
	authCollection  *mongo.Collection
	ridesCollection *mongo.Collection
	returnAfter     options.ReturnDocument
	upsert          = true
)

// Init - Initialize services
func Init() {
	returnAfter = options.After
	userCollection = db.GetCollection("users")
	ridesCollection = db.GetCollection("rides")
	authCollection = db.GetCollection("auth")
}
