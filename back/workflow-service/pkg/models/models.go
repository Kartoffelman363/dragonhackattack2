package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// when changing file make sure to update across repo

type Variable struct {
	Id      string `json:"id" bson:"id"`
	VarName string `json:"varname" bson:"varname"`
	Type    string `json:"type" bson:"type"`
	Value   string `json:"value" bson:"value"`
}

type Block struct {
	ID              string     `json:"id" bson:"id"`
	InputVariables  []Variable `json:"inputvariables" bson:"inputvariables"`
	OutputVariables []Variable `json:"outputvariables" bson:"outputvariables"`
	Code            string     `json:"code" bson:"code"`
}

type Workflow struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id"`
	Metadata         string             `json:"metadata" bson:"metadata"`
	Blocks           []Block            `json:"blocks" bson:"blocks"`
	InitialVariables []Variable         `json:"initialvariables" bson:"initialvariables"`
}

type WorkflowCreate struct {
	Metadata         string     `json:"metadata" bson:"metadata"`
	Blocks           []Block    `json:"blocks" bson:"blocks"`
	InitialVariables []Variable `json:"initialvariables" bson:"initialvariables"`
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
