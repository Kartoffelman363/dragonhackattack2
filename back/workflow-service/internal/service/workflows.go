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

	collection := mongodb.GetCollection("yourDatabase", "workflows")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []models.Workflow

	fmt.Print(results)
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

	collection := mongodb.GetCollection("yourDatabase", "workflows")
	var result models.Workflow
	if err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func DeleteWorkflow(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := mongodb.GetCollection("yourDatabase", "workflows")
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
		return nil, fmt.Errorf("could not insert workflows")
	}

	var insertedWorkflow models.Workflow

	if err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&insertedWorkflow); err != nil {
		return nil, err
	}

	return &insertedWorkflow, nil
}

func UpdateWorkflow(id primitive.ObjectID, updateData models.WorkflowCreate) (*models.Workflow, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := mongodb.GetCollection("yourDatabase", "workflows")

	update := bson.M{
		"$set": updateData,
	}

	updateResult, err := collection.UpdateByID(ctx, id, update)
	if err != nil {
		return nil, fmt.Errorf("could not update workflow: %v", err)
	}

	if updateResult.MatchedCount == 0 {
		return nil, fmt.Errorf("no workflow found with ID %v", id)
	}

	var updatedWorkflow models.Workflow
	if err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&updatedWorkflow); err != nil {
		return nil, fmt.Errorf("could not retrieve updated workflow: %v", err)
	}

	return &updatedWorkflow, nil
}
