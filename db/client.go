package db

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/Glitchfix/golagobar/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	db     *mongo.Database
)

// Init - initialize database
func Init(ctx context.Context) {
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d/?authSource=%s",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.AuthSource,
	)
	logrus.Infoln(uri)
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(opts)
	if nil != err {
		logrus.Fatalln(err)
		return
	}
	err = client.Connect(ctx)
	if nil != err {
		logrus.Fatalln(err)
		return
	}

	if err = client.Ping(ctx, opts.ReadPreference); nil != err {
		logrus.Infoln("could not ping to mongo db service: %v\n", err)
		return
	}

	db = client.Database(config.DB.Database)

	logrus.Infoln("Connected to database")
}

// GetCollection - Get a collection by it's name
func GetCollection(collectionName string) *mongo.Collection {
	return db.Collection(collectionName)
}
