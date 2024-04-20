package service

import (
	"context"
	"document-service/internal/mongodb"
	models "document-service/pkg/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllDocuments() (*models.Documents, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := mongodb.GetCollection("yourDatabase", "documents")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []models.Document
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	var response = models.Documents{
		Documents: results,
	}
	return &response, nil
}

func GetDocumentByID(id primitive.ObjectID) (*models.Document, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := mongodb.GetCollection("yourDatabase", "documents")
	var result models.Document
	if err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func DeleteDocument(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := mongodb.GetCollection("yourDatabase", "documents")
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func CreateDocument(document models.DocumentCreate) (*models.Document, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := mongodb.GetCollection("yourDatabase", "documents")

	result, err := collection.InsertOne(ctx, document)
	if err != nil {
		return nil, err
	}

	objID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("could not insert workflow")
	}

	var insertedDocument models.Document

	if err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&insertedDocument); err != nil {
		return nil, err
	}

	return &insertedDocument, nil
}

func UpdateDocument(id primitive.ObjectID, updateData models.DocumentCreate) (*models.Document, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := mongodb.GetCollection("yourDatabase", "documents")

	update := bson.M{
		"$set": updateData,
	}

	updateResult, err := collection.UpdateByID(ctx, id, update)
	if err != nil {
		return nil, fmt.Errorf("could not update document: %v", err)
	}

	if updateResult.MatchedCount == 0 {
		return nil, fmt.Errorf("no workflow found with ID %v", id)
	}

	var updatedDocument models.Document
	if err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&updatedDocument); err != nil {
		return nil, fmt.Errorf("could not retrieve updated workflow: %v", err)
	}

	return &updatedDocument, nil
}
