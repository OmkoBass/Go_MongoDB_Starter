package database

import (
	"Go_Fiber_Starter/config"
	Models "Go_Fiber_Starter/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbName = config.Config("MongoDBDATABASE")
var dbUsername = config.Config("MongoDBUSERNAME")
var dbPassword = config.Config("MongoDBPASSWORD")

var mongoURI = "mongodb+srv://" + dbUsername + ":" + dbPassword + "@practice.xcm7r.mongodb.net/" + dbName + "?retryWrites=true&w=majority"

var mg Models.MongoInstance

func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	mg = Models.MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}

func GetDB() Models.MongoInstance {
    return mg
}
