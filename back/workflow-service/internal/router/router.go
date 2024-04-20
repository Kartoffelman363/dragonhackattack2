package router

import (
	"workflow-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/workflows", handlers.GetAllWorkflows)
	r.GET("/workflows/:id", handlers.GetWorkflowByID)
	r.DELETE("/workflows/:id", handlers.DeleteWorkflow)
	r.POST("/workflows", handlers.CreateWorkflow)
	r.POST("/workflows/:id", handlers.UpdateWorkflow)

	return r
}
