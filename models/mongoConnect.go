//Package models 用來define model
package models

import (
	"context"

	"iBP/helper"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connect connects to MongoDB
func Connect() *mongo.Client {
	var err error
	var cred options.Credential
	// 連docker container
	cred.Username = viper.GetString("mongo.user")
	cred.Password = viper.GetString("mongo.pass")
	URI := "mongodb://" + viper.GetString("mongo.host") + ":" + viper.GetString("mongo.port")
	clientOptions := options.Client().ApplyURI(URI).SetAuth(cred)
	// 連azure cosmos
	// mongoDBConnectionString := viper.GetString("mongo.azure_cosmos")
	// clientOptions := options.Client().ApplyURI(mongoDBConnectionString).SetDirect(true)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	c, _ := mongo.NewClient(clientOptions)

	err = c.Connect(ctx)

	if err != nil {
		helper.Fatal("unable to initialize connection" + err.Error())
	}
	err = c.Ping(ctx, readpref.Primary())
	if err != nil {
		helper.Fatal("unable to connect " + err.Error())
	}
	return c
}
