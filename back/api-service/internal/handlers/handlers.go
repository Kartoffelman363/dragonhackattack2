package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	models "api-service/pkg/models"

	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, statusCode int, err string) {
	c.JSON(statusCode, gin.H{"error": err})
}

func makeHTTPRequest(method, url string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server responded with status code: %d", response.StatusCode)
	}
	return response, nil
}

func GetAllDocuments(c *gin.Context) {
	docUrl := os.Getenv("DOC_URL")
	if docUrl == "" {
		respondWithError(c, 500, "DOC_URL environment variable is not set")
		return
	}
	url := docUrl + "/documents/"
	response, err := makeHTTPRequest("GET", url, nil)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		respondWithError(c, 500, "Failed to read response body: "+err.Error())
		return
	}

	var documents models.Documents
	if err = json.Unmarshal(body, &documents); err != nil {
		respondWithError(c, 500, "Failed to unmarshal response: "+err.Error())
		return
	}
	c.JSON(200, documents)
}

func GetDocumentByID(c *gin.Context) {
	docUrl := os.Getenv("DOC_URL")
	if docUrl == "" {
		respondWithError(c, 500, "DOC_URL environment variable is not set")
		return
	}
	url := docUrl + "/documents/" + c.Param("id")
	response, err := makeHTTPRequest("GET", url, nil)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		respondWithError(c, 500, "Failed to read response body: "+err.Error())
		return
	}

	var document models.Document
	if err = json.Unmarshal(body, &document); err != nil {
		respondWithError(c, 500, "Failed to unmarshal response: "+err.Error())
		return
	}
	c.JSON(200, document)
}

func DeleteDocument(c *gin.Context) {
	docUrl := os.Getenv("DOC_URL")
	if docUrl == "" {
		respondWithError(c, 500, "DOC_URL environment variable is not set")
		return
	}
	url := docUrl + "/documents/" + c.Param("id")
	response, err := makeHTTPRequest("DELETE", url, nil)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	response.Body.Close()
	c.JSON(200, gin.H{"message": "Document deleted successfully"})
}

func CreateDocument(c *gin.Context) {
	docUrl := os.Getenv("DOC_URL")
	if docUrl == "" {
		respondWithError(c, 500, "DOC_URL environment variable is not set")
		return
	}
	url := docUrl + "/documents"
	response, err := makeHTTPRequest("POST", url, c.Request.Body)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		respondWithError(c, 500, "Failed to read response body: "+err.Error())
		return
	}

	var document models.Document
	if err = json.Unmarshal(body, &document); err != nil {
		respondWithError(c, 500, "Failed to unmarshal response: "+err.Error())
		return
	}
	c.JSON(200, document)
}

func GetAllWorkflows(c *gin.Context) {
	workflowUrl := os.Getenv("WORKFLOW_URL")
	if workflowUrl == "" {
		respondWithError(c, 500, "WORKFLOW_URL environment variable is not set")
		return
	}
	url := workflowUrl + "/workflows/"
	response, err := makeHTTPRequest("GET", url, nil)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		respondWithError(c, 500, "Failed to read response body: "+err.Error())
		return
	}

	var workflows models.Workflows
	if err = json.Unmarshal(body, &workflows); err != nil {
		respondWithError(c, 500, "Failed to unmarshal response: "+err.Error())
		return
	}
	c.JSON(200, workflows)
}

func GetWorkflowByID(c *gin.Context) {
	workflowUrl := os.Getenv("WORKFLOW_URL")
	if workflowUrl == "" {
		respondWithError(c, 500, "WORKFLOW_URL environment variable is not set")
		return
	}
	url := workflowUrl + "/workflows/" + c.Param("id")
	response, err := makeHTTPRequest("GET", url, nil)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		respondWithError(c, 500, "Failed to read response body: "+err.Error())
		return
	}

	var workflow models.Workflow
	if err = json.Unmarshal(body, &workflow); err != nil {
		respondWithError(c, 500, "Failed to unmarshal response: "+err.Error())
		return
	}
	c.JSON(200, workflow)
}

func DeleteWorkflow(c *gin.Context) {
	workflowUrl := os.Getenv("WORKFLOW_URL")
	if workflowUrl == "" {
		respondWithError(c, 500, "WORKFLOW_URL environment variable is not set")
		return
	}
	url := workflowUrl + "/workflows/" + c.Param("id")
	response, err := makeHTTPRequest("DELETE", url, nil)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	response.Body.Close()
	c.JSON(200, gin.H{"message": "Workflow deleted successfully"})
}

func CreateWorkflow(c *gin.Context) {
	workflowUrl := os.Getenv("WORKFLOW_URL")
	if workflowUrl == "" {
		respondWithError(c, 500, "WORKFLOW_URL environment variable is not set")
		return
	}
	url := workflowUrl + "/workflows"
	response, err := makeHTTPRequest("POST", url, c.Request.Body)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		respondWithError(c, 500, "Failed to read response body: "+err.Error())
		return
	}

	var workflow models.Workflow
	if err = json.Unmarshal(body, &workflow); err != nil {
		respondWithError(c, 500, "Failed to unmarshal response: "+err.Error())
		return
	}
	c.JSON(200, workflow)
}

func RunWorkflow(c *gin.Context) {
	workflowUrl := os.Getenv("WORKFLOW_URL")
	if workflowUrl == "" {
		respondWithError(c, 500, "WORKFLOW_URL environment variable is not set")
		return
	}

	url := workflowUrl + "/workflows/" + c.Param("id")
	response, err := makeHTTPRequest("GET", url, nil)
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		respondWithError(c, 500, "Failed to read response body: "+err.Error())
		return
	}

	coderunnerUrl := os.Getenv("CODERUNNER_URL")
	url = coderunnerUrl + "/workflows/" + c.Param("id") + "/run"
	runResponse, err := makeHTTPRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		respondWithError(c, 500, err.Error())
		return
	}
	runResponse.Body.Close()

	c.JSON(200, gin.H{"message": "Workflow run successfully"})
}
