package handlers

import (
	"encoding/json"

	models "api-service/pkg/models"

	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{"error": err.Error()})
}

func GetAllDocuments(c *gin.Context) {
	req := models.MessageRequest{
		Operation: "get_all",
	}
	res, err := sendRabbitMQRequest("document_tasks", req)
	if err != nil {
		respondWithError(c, 500, err)
		return
	}
	if res == nil {
		c.JSON(404, gin.H{"error": "No documents found"})
		return
	}
	body := models.Documents{}
	err = json.Unmarshal(res, &body)
	if err != nil {
		respondWithError(c, 500, err)
		return
	}
	c.JSON(200, body)
}

func GetDocumentByID(c *gin.Context) {
	req := models.MessageRequest{
		Operation: "get_by_id",
		Payload: models.DocumentGetByIDRequest{
			ID: c.Param("id"),
		},
	}
	res, err := sendRabbitMQRequest("document_tasks", req)
	if err != nil {
		respondWithError(c, 500, err)
		return
	}
	if res == nil {
		c.JSON(404, gin.H{"error": "No documents found"})
		return
	}
	body := models.Document{}
	err = json.Unmarshal(res, &body)
	if err != nil {
		respondWithError(c, 500, err)
		return
	}
	c.JSON(200, body)
}

func DeleteDocument(c *gin.Context) {
	req := models.MessageRequest{
		Operation: "delete",
		Payload: models.DocumentDeletionRequest{
			ID: c.Param("id"),
		},
	}
	_, err := sendRabbitMQRequest("document_tasks", req)
	if err != nil {
		respondWithError(c, 500, err)
		return
	}
	c.JSON(200, "")
}

// TOOD: CreateDocument request + struct

func GetAllWorkflows(c *gin.Context) {
	req := models.MessageRequest{
		Operation: "get_all",
	}
	res, err := sendRabbitMQRequest("workflow_tasks", req)
	if err != nil {
		respondWithError(c, 500, err)
		return
	}
	if res == nil {
		c.JSON(404, gin.H{"error": "No documents found"})
		return
	}
	body := models.Workflows{}
	err = json.Unmarshal(res, &body)
	if err != nil {
		respondWithError(c, 500, err)
		return
	}
	c.JSON(200, body)
}

func GetWorkflowByID(c *gin.Context) {
	req := models.MessageRequest{
		Operation: "get_by_id",
		Payload: models.WorkflowGetByIDRequest{
			ID: c.Param("id"),
		},
	}
	res, err := sendRabbitMQRequest("workflow_tasks", req)
	if err != nil {
		respondWithError(c, 500, err)
		return
	}
	if res == nil {
		c.JSON(404, gin.H{"error": "No documents found"})
		return
	}
	body := models.Workflow{}
	err = json.Unmarshal(res, &body)
	if err != nil {
		respondWithError(c, 500, err)
		return
	}
	c.JSON(200, body)
}

func DeleteWorkflow(c *gin.Context) {
	req := models.MessageRequest{
		Operation: "delete",
		Payload: models.WorkflowDeletionRequest{
			ID: c.Param("id"),
		},
	}
	_, err := sendRabbitMQRequest("workflow_tasks", req)
	if err != nil {
		respondWithError(c, 500, err)
		return
	}
	c.JSON(200, "")
}

// TOOD: CreateWorkflow request + struct

func RunWorkflow(c *gin.Context) {
	reqWorkflow := models.MessageRequest{
		Operation: "get_by_id",
		Payload: models.WorkflowGetByIDRequest{
			ID: c.Param("id"),
		},
	}
	res, err := sendRabbitMQRequest("workflow_tasks", reqWorkflow)
	if err != nil {
		respondWithError(c, 500, err)
		return
	}
	if res == nil {
		c.JSON(404, gin.H{"error": "No documents found"})
		return
	}
	body := models.Workflow{}
	err = json.Unmarshal(res, &body)
	if err != nil {
		respondWithError(c, 500, err)
		return
	}

	// TODO: unmarshal and send the workflow to the runner!
	reqRunner := models.MessageRequest{
		Operation: "run",
		Payload:   body,
	}
	sendRabbitMQRequest("coderunner_tasks", reqRunner)
}
