package storage

import (
	"context"
	"fmt"
	. "fofo/common"
	"fofo/config"
	"fofo/entity"
	"fofo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var MongoClient *mongo.Client
var ServiceModel *mongo.Collection
var ctx context.Context

func init() {
	var err error
	mongoAddr := fmt.Sprintf("mongodb://%s:%d", config.Config.GetString("MongoDBHost"), config.Config.GetInt("MongoDBPort"))
	Logger.Info("mongoAddr = ", mongoAddr)
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	MongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoAddr))
	if err != nil {
		Logger.Error(err)
		return
	}
	err = MongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		Logger.Error(err)
		return
	}
	ServiceModel = MongoClient.Database("fofo").Collection("service")
}

func AddService(service *entity.Service) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	data := utils.Struct2Map(*service)
	_, err := ServiceModel.InsertOne(ctx, bson.M(data))
	if err != nil {
		Logger.Error(err)
		return err
	}
	return nil
}

func GetService(serviceID string) (*entity.Service, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"ID": serviceID}
	var service entity.Service
	err := ServiceModel.FindOne(ctx, filter).Decode(&service)
	if err != nil {
		Logger.Error(err)
		return nil, err
	}
	return &service, nil
}

func UpdateServiceTime(serviceID string, healthCheckTime int64) (int64, error) {
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.D{{"ID", serviceID}}
	data := bson.D{
		{
			"$set", bson.D{{"HealthCheckTime", healthCheckTime}},
		},
	}
	r, err := ServiceModel.UpdateOne(ctx, filter, data)
	if err != nil {
		Logger.Error(err)
		return 0, err
	}
	return r.ModifiedCount, nil
}

// SearchService 搜索服务
func SearchService(name string, group string) ([]*entity.Service, error) {
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	cond := make(map[string]interface{})
	if name != "" {
		cond["Name"] = name
	}
	if group != "" {
		cond["Group"] = group
	}
	filter := bson.M(cond)
	cursor, err := ServiceModel.Find(ctx, filter)
	if err != nil {
		Logger.Error(err)
		return nil, err
	}
	var result []*entity.Service
	for cursor.Next(context.TODO()) {
		var service entity.Service
		err = cursor.Decode(&service)
		if err != nil {
			Logger.Error(err)
			return nil, err
		}
		if !service.IsValid() {
			continue
		}
		result = append(result, &service)
	}
	err = cursor.Close(ctx)
	if err != nil {
		Logger.Error(err)
	}
	return result, nil
}
