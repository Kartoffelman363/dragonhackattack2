package router

import (
	"api-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(CORSMiddleware())

	r.GET("/documents", handlers.GetAllDocuments)
	r.GET("/documents/:id", handlers.GetDocumentByID)
	r.DELETE("/documents/:id", handlers.DeleteDocument)
	r.POST("/documents", handlers.CreateDocument)

	r.GET("/workflows", handlers.GetAllWorkflows)
	r.GET("/workflows/:id", handlers.GetWorkflowByID)
	r.DELETE("/workflows/:id", handlers.DeleteWorkflow)
	r.POST("/workflows", handlers.CreateWorkflow)
	r.POST("/workflows/:id/run", handlers.RunWorkflow)

	return r
}
