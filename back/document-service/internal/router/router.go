package router

import (
	"document-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/documents", handlers.GetAllDocuments)
	r.GET("/documents/:id", handlers.GetDocumentByID)
	r.DELETE("/documents/:id", handlers.DeleteDocument)

	return r
}
