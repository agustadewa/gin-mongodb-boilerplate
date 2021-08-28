package repository

import "go.mongodb.org/mongo-driver/mongo"

// MongoRepo is a repo object
type MongoRepo struct {
	// mongo collection, if u want a mongo client, just jump into its mongo.Database then mongo.Client
	coll *mongo.Collection
}

// New generates new repo object
func New(client *mongo.Client) *MongoRepo {
	// get mongo db config
	dbName := ""
	collName := ""

	return &MongoRepo{coll: client.Database(dbName).Collection(collName)}
}
