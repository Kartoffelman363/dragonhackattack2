package router

import (
	"document-service/internal/handlers"

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
	r.POST("/documents/:id", handlers.UpdateDocument)

	return r
}
