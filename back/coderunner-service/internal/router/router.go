package router

import (
	"coderunner-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/workflows/:id/run", handlers.RunWorkflow)

	return r
}
