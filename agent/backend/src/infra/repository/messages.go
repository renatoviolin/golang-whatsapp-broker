package repository

import (
	"context"
	"errors"
	"whatsapp-client/dto"

	"go.mongodb.org/mongo-driver/bson"
)

type MessageRepository struct {
	mongoClient *MongoClient
	collection  string
}

func NewMessageRepository(client *MongoClient, collection string) MessageRepository {
	return MessageRepository{
		mongoClient: client,
		collection:  collection,
	}
}

func (l *MessageRepository) FindAll(waID string) ([]dto.MessageRepository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), l.mongoClient.Timeout)
	defer cancel()

	collection := l.mongoClient.Client.Database(l.mongoClient.Database).Collection(l.collection)
	filter := bson.M{"$and": []bson.M{
		{"waID": waID},
		{"isRead": false},
	}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var messages []dto.MessageRepository
	if err = cursor.All(context.TODO(), &messages); err != nil {
		return nil, err
	}

	return messages, nil
}

func (l *MessageRepository) Create(input dto.MessageRepository) error {
	ctx, cancel := context.WithTimeout(context.Background(), l.mongoClient.Timeout)
	defer cancel()

	collection := l.mongoClient.Client.Database(l.mongoClient.Database).Collection(l.collection)
	_, err := collection.InsertOne(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (l *MessageRepository) MarkAsRead(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), l.mongoClient.Timeout)
	defer cancel()

	collection := l.mongoClient.Client.Database(l.mongoClient.Database).Collection(l.collection)
	filter := bson.M{"id": id}
	update := bson.M{"$set": bson.M{"isRead": true}}
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if res.ModifiedCount == 0 {
		return errors.New(id + " not found to be updated")
	}
	return nil
}
