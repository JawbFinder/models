package models

import "go.mongodb.org/mongo-driver/mongo"

type MongoDatastore struct {
	Session *mongo.Client
}
