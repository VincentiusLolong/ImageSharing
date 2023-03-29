package database

import (
	"context"
	"fmt"
	"log"
	"mestorage/conf/env"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Dbs interface {
	MongoImage() *mongo.Collection
	MongoAccount() *mongo.Collection
}

type database struct {
	Collection *mongo.Collection
}

func New() Dbs {
	return &database{}
}

func MongoConnect() *mongo.Client {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(env.Env("MONGO_API")).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

var mongose *mongo.Client = MongoConnect()
var mdb string = env.Env("MONGO_DATABASE")

func (db *database) MongoImage() *mongo.Collection {
	db.Collection = mongose.Database(mdb).Collection(env.Env("MONGO_2"))
	return db.Collection
}

func (db *database) MongoAccount() *mongo.Collection {
	db.Collection = mongose.Database(mdb).Collection(env.Env("MONGO_1"))
	return db.Collection
}
