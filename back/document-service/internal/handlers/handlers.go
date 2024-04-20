package handlers

import (
	service "document-service/internal/service"
	models "document-service/pkg/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func respondWithError(c *gin.Context, statusCode int, err string) {
	c.JSON(statusCode, gin.H{"error": err})
}

func GetAllDocuments(c *gin.Context) {

	res, err := service.GetAllDocuments()
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	if res == nil {
		respondWithError(c, 404, "No documents found")
		return
	}
	body := models.Documents{}
	err = bson.Unmarshal(res, &body)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	c.JSON(200, body)
}

func GetDocumentByID(c *gin.Context) {
	req := models.DocumentGetByIDRequest{
		ID: c.Param("id"),
	}

	out, err := service.GetDocumentByID(req.ID)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}

	body := models.Document{}
	err = bson.Unmarshal(out, &body)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	c.JSON(200, body)
}

func DeleteDocument(c *gin.Context) {
	req := models.DocumentDeletionRequest{
		ID: c.Param("id"),
	}
	err := service.DeleteDocument(req.ID)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	c.JSON(200, "")
}
