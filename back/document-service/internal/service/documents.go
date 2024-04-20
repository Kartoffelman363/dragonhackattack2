package service

import (
	"context"
	"document-service/internal/mongodb"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllDocuments() ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := mongodb.GetCollection("yourDatabase", "documents")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	bsonBytes, err := bson.Marshal(results)
	if err != nil {
		return nil, err
	}
	return bsonBytes, nil
}

func GetDocumentByID(id string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}

	collection := mongodb.GetCollection("yourDatabase", "documents")
	var result bson.M
	if err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&result); err != nil {
		return nil, err
	}
	bsonBytes, err := bson.Marshal(result)
	if err != nil {
		return nil, err
	}
	return bsonBytes, nil
}

func DeleteDocument(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid ID format: %v", err)
	}

	collection := mongodb.GetCollection("yourDatabase", "documents")
	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
