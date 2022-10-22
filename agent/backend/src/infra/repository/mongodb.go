package repository

import (
	"context"
	"time"
	"whatsapp-client/infra/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClient struct {
	Client   mongo.Client
	Timeout  time.Duration
	Database string
}

func NewMongoClient(mongoURL string, database string, timeout int) (*MongoClient, error) {
	mongoClient := &MongoClient{
		Timeout:  time.Duration(timeout) * time.Second,
		Database: database,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		logger.Fatal("new-mongo-client", err.Error())
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logger.Fatal("new-mongo-client", err.Error())
	}

	mongoClient.Client = *client
	return mongoClient, nil
}
