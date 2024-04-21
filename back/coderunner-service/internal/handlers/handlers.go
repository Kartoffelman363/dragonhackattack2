package handlers

import (
	"coderunner-service/internal/service"
	models "coderunner-service/pkg/models"
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, statusCode int, err string) {
	c.JSON(statusCode, gin.H{"error": err})
}

func RunWorkflow(c *gin.Context) {
	reqWorkflow := models.Workflow{}

	jsonData, err := io.ReadAll(c.Request.Body)

	if err != nil {
		respondWithError(c, 500, err.Error())
	}
	err = json.Unmarshal(jsonData, &reqWorkflow)

	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}

	value, err := service.StartParsing(reqWorkflow)

	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}

	c.JSON(200, value)
}
