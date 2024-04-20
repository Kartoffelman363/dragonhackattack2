package handlers

import (
	service "workflow-service/internal/service"
	models "workflow-service/pkg/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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
	body := models.Workflows{}
	err = bson.Unmarshal(res, &body)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	c.JSON(200, body)
}

func GetWorkflowByID(c *gin.Context) {
	req := models.WorkflowGetByIDRequest{
		ID: c.Param("id"),
	}
	out, err := service.GetWorkflowByID(req.ID)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}

	body := models.Workflow{}
	err = bson.Unmarshal(out, &body)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	c.JSON(200, body)
}

func DeleteWorkflow(c *gin.Context) {
	req := models.WorkflowDeletionRequest{
		ID: c.Param("id"),
	}
	err := service.DeleteWorkflow(req.ID)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	c.JSON(200, "")
}
