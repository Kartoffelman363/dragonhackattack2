package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	models "api-service/pkg/models"

	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, statusCode int, err string) {
	c.JSON(statusCode, gin.H{"error": err})
}

func GetAllDocuments(c *gin.Context) {
	docUrl := os.Getenv("DOC_URL")
	if docUrl == "" {
		respondWithError(c, 500, "internal server error")
	}
	url := docUrl + "/documents/"
	res, err := http.NewRequest("GET", url, nil)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	if res == nil {
		respondWithError(c, 404, "no documents found")
		return
	}
	defer res.Body.Close()
	text, err := io.ReadAll(res.Body)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	body := models.Documents{}
	err = json.Unmarshal(text, &body)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	c.JSON(200, body)
}

func GetDocumentByID(c *gin.Context) {
	docUrl := os.Getenv("DOC_URL")
	if docUrl == "" {
		respondWithError(c, 500, "internal server error")
	}

	url := docUrl + "/documents/" + c.Param("id")
	res, err := http.NewRequest("GET", url, nil)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	if res == nil {
		respondWithError(c, 404, "no documents found")
		return
	}
	defer res.Body.Close()
	text, err := io.ReadAll(res.Body)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	body := models.Document{}
	err = json.Unmarshal(text, &body)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	c.JSON(200, body)
}

func DeleteDocument(c *gin.Context) {
	docUrl := os.Getenv("DOC_URL")
	if docUrl == "" {
		respondWithError(c, 500, "internal server error")
	}

	url := docUrl + "/documents/" + c.Param("id")
	_, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	c.JSON(200, "")
}

func CreateDocument(c *gin.Context) {
	docUrl := os.Getenv("DOC_URL")
	if docUrl == "" {
		respondWithError(c, 500, "internal server error")
	}
	url := docUrl + "/documents"
	res, err := http.NewRequest("POST", url, c.Request.Body)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	if res == nil {
		respondWithError(c, 500, "document creation failed")
		return
	}
	defer res.Body.Close()
	text, err := io.ReadAll(res.Body)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	body := models.Document{}
	err = json.Unmarshal(text, &body)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	c.JSON(200, body)
}

// TOOD: CreateDocument request + struct

func GetAllWorkflows(c *gin.Context) {
	workflowUrl := os.Getenv("WORKFLOW_URL")
	if workflowUrl == "" {
		respondWithError(c, 500, "internal server error")
	}
	url := workflowUrl + "/workflows/"
	res, err := http.NewRequest("GET", url, nil)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	if res == nil {
		respondWithError(c, 404, "no workflows found")
		return
	}
	defer res.Body.Close()
	text, err := io.ReadAll(res.Body)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	body := models.Workflows{}
	err = json.Unmarshal(text, &body)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	c.JSON(200, body)
}

func GetWorkflowByID(c *gin.Context) {
	workflowUrl := os.Getenv("WORKFLOW_URL")
	if workflowUrl == "" {
		respondWithError(c, 500, "internal server error")
	}

	url := workflowUrl + "/workflows/" + c.Param("id")
	res, err := http.NewRequest("GET", url, nil)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	if res == nil {
		respondWithError(c, 404, "no workflows found")
		return
	}
	defer res.Body.Close()
	text, err := io.ReadAll(res.Body)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	body := models.Workflow{}
	err = json.Unmarshal(text, &body)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	c.JSON(200, body)
}

func DeleteWorkflow(c *gin.Context) {
	workflowUrl := os.Getenv("WORKFLOW_URL")
	if workflowUrl == "" {
		respondWithError(c, 500, "internal server error")
	}

	url := workflowUrl + "/workflows/" + c.Param("id")
	_, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	c.JSON(200, "")
}

func CreateWorkflow(c *gin.Context) {
	workflowUrl := os.Getenv("WORKFLOW_URL")
	if workflowUrl == "" {
		respondWithError(c, 500, "internal server error")
	}
	url := workflowUrl + "/workflows"
	res, err := http.NewRequest("POST", url, c.Request.Body)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	if res == nil {
		respondWithError(c, 500, "workflow creation failed")
		return
	}
	defer res.Body.Close()
	text, err := io.ReadAll(res.Body)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	body := models.Workflow{}
	err = json.Unmarshal(text, &body)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	c.JSON(200, body)
}

func RunWorkflow(c *gin.Context) {
	workflowUrl := os.Getenv("WORKFLOW_URL")
	if workflowUrl == "" {
		respondWithError(c, 500, "internal server error")
	}

	url := workflowUrl + "/workflows/" + c.Param("id")
	res, err := http.NewRequest("GET", url, nil)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	if res == nil {
		respondWithError(c, 404, "workflow doesn't exist")
		return
	}
	defer res.Body.Close()
	text, err := io.ReadAll(res.Body)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	coderunnerUrl := os.Getenv("CODERUNNER_URL")
	url = coderunnerUrl + "/workflows/" + c.Param("id") + "/run"
	_, err = http.NewRequest("POST", url, bytes.NewBuffer(text))
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}

	c.JSON(200, "")
}
