package global

import "go.mongodb.org/mongo-driver/mongo"

var (
	MongoClientYD *mongo.Client
	UserColl      *mongo.Collection
	RoleColl      *mongo.Collection
	ApiColl       *mongo.Collection
)
