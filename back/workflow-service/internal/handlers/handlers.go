package handlers

import (
	service "workflow-service/internal/service"
	models "workflow-service/pkg/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func respondWithError(c *gin.Context, statusCode int, err string) {
	c.JSON(statusCode, gin.H{"error": err})
}

func GetAllWorkflows(c *gin.Context) {
	res, err := service.GetAllWorflows()
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	if res == nil {
		respondWithError(c, 404, "No documents found")
		return
	}
	c.JSON(200, *res)
}

func GetWorkflowByID(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	res, err := service.GetWorkflowByID(id)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	c.JSON(200, *res)
}

func DeleteWorkflow(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	err = service.DeleteWorkflow(id)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	c.JSON(200, "")
}

func CreateWorkflow(c *gin.Context) {
	var newWorkflow models.WorkflowCreate
	if err := c.BindJSON(&newWorkflow); err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	res, err := service.CreateWorkflow(newWorkflow)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	c.JSON(200, *res)
}

func UpdateWorkflow(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	var workflow models.WorkflowCreate
	if err := c.BindJSON(&workflow); err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	res, err := service.UpdateWorkflow(id, workflow)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	c.JSON(200, *res)
}
