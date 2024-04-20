package handlers

import (
	service "document-service/internal/service"
	models "document-service/pkg/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	c.JSON(200, *res)
}

func GetDocumentByID(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	res, err := service.GetDocumentByID(id)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	c.JSON(200, *res)
}

func DeleteDocument(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	err = service.DeleteDocument(id)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	c.JSON(200, "")
}

func CreateDocument(c *gin.Context) {
	var newDocument models.DocumentCreate
	if err := c.BindJSON(&newDocument); err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	res, err := service.CreateDocument(newDocument)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	c.JSON(200, *res)
}
