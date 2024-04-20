package handlers

import (
	models "coderunner-service/pkg/models"
	_ "fmt"

	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, statusCode int, err string) {
	c.JSON(statusCode, gin.H{"error": err})
}

func RunWorkflow(c *gin.Context) {
	reqWorkflow := models.Workflow{
		ID:   c.Param("id"),
		Name: c.Param("name"),
	}
	_ = reqWorkflow

}
