package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// when changing file make sure to update across repo

type Status struct {
	StatusCode string `json:"statuscode" bson:"statuscode"`
}

type Variable struct {
	VarName string `json:"varname" bson:"varname"`
	Type    string `json:"type" bson:"type"`
	Value   string `json:"value" bson:"value"`
	Example string `json:"example" bson:"example"`
}

type MetaData struct {
	X int `json:"x" bson:"x"`
	Y int `json:"y" bson:"y"`
}

type Workflow struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	Name            string             `json:"name" bson:"name"`
	Metadata        MetaData           `json:"metadata" bson:"metadata"`
	InputVariables  []Variable         `json:"input_variables" bson:"input_variables"`
	OutputVariables []Variable         `json:"output_variables" bson:"output_variables"`
	Code            string             `json:"code" bson:"code"`
	Blocks          Workflows          `json:"blocks" bson:"blocks"`
}

type WorkflowCreate struct {
	Name            string     `json:"name" bson:"name"`
	Metadata        MetaData   `json:"metadata" bson:"metadata"`
	InputVariables  []Variable `json:"input_variables" bson:"input_variables"`
	OutputVariables []Variable `json:"output_variables" bson:"output_variables"`
	Code            string     `json:"code" bson:"code"`
	Blocks          Workflows  `json:"blocks" bson:"blocks"`
}

type Document struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Content   string             `json:"content" bson:"content"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

type DocumentCreate struct {
	Name      string    `json:"name" bson:"name"`
	Content   string    `json:"content" bson:"content"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

type Workflows struct {
	Workflows []Workflow `json:"workflows" bson:"workflows"`
}

type Documents struct {
	Documents []Document `json:"documents" bson:"documents"`
}
