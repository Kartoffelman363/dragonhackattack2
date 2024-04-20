package service

import (
	"context"
	"fmt"
	"time"
	"workflow-service/internal/mongodb"
	models "workflow-service/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllWorflows() (*models.Workflows, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := mongodb.GetCollection("yourDatabase", "worflows")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []models.Workflow
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	var response = models.Workflows{
		Workflows: results,
	}
	return &response, nil
}

func GetWorkflowByID(id primitive.ObjectID) (*models.Workflow, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := mongodb.GetCollection("yourDatabase", "worflows")
	var result models.Workflow
	if err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func DeleteWorkflow(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := mongodb.GetCollection("yourDatabase", "worflows")
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func CreateWorkflow(workflow models.WorkflowCreate) (*models.Workflow, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := mongodb.GetCollection("yourDatabase", "workflows")

	result, err := collection.InsertOne(ctx, workflow)
	if err != nil {
		return nil, err
	}

	objID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("could not insert workflow")
	}

	var insertedWorkflow models.Workflow

	if err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&insertedWorkflow); err != nil {
		return nil, err
	}

	return &insertedWorkflow, nil
}
