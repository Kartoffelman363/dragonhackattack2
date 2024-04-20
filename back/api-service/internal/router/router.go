package router

import (
	"api-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/documents", handlers.GetAllDocuments)
	r.GET("/documents/:id", handlers.GetDocumentByID)
	r.DELETE("/documents/:id", handlers.DeleteDocument)

	r.GET("/workflows", handlers.GetAllWorkflows)
	r.GET("/workflows/:id", handlers.GetWorkflowByID)
	r.DELETE("/workflows/:id", handlers.DeleteWorkflow)
	r.POST("/workflows/:id/run", handlers.RunWorkflow)

	return r
}
