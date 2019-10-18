package storage

import (
	"context"
	. "fofo/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var MongoClient *mongo.Client
var ServiceModel *mongo.Collection

func init() {
	var err error
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	MongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		Logger.Error(err)
		return
	}
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = MongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		Logger.Error(err)
		return
	}
	ServiceModel = MongoClient.Database("fofo").Collection("service")
}

func GetService(name string, condition map[string]interface{}) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M(condition)
	ServiceModel.FindOne(ctx, filter)
}

func AddService() {

}
