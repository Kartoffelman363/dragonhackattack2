package internalhandlers

import (
	"bytes"
	models "coderunner-service/pkg/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetDocumentByID(id string) (*models.Document, error) {
	docUrl := os.Getenv("DOC_URL")
	if docUrl == "" {
		return nil, fmt.Errorf("internal server error")
	}

	url := docUrl + "/documents/" + id
	res, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no documents found")
	}
	defer res.Body.Close()

	text, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	body := models.Document{}
	err = json.Unmarshal(text, &body)
	if err != nil {
		return nil, err
	}

	return &body, nil
}

func DeleteDocument(id string) error {
	docUrl := os.Getenv("DOC_URL")
	if docUrl == "" {
		return fmt.Errorf("internal server error")
	}

	url := docUrl + "/documents/" + id
	_, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return nil
}

func CreateDocument(input *models.DocumentCreate) (*models.Document, error) {
	docUrl := os.Getenv("DOC_URL")
	if docUrl == "" {
		return nil, fmt.Errorf("internal server error")
	}

	byteArray, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	url := docUrl + "/documents"
	res, err := http.NewRequest("POST", url, bytes.NewBuffer(byteArray))
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("document creation failed")
	}
	defer res.Body.Close()

	text, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	body := models.Document{}
	err = json.Unmarshal(text, &body)
	if err != nil {
		return nil, err
	}

	return &body, nil
}

func GetWorkflowByID(id string) (*models.Workflow, error) {
	workflowUrl := os.Getenv("WORKFLOW_URL")
	if workflowUrl == "" {
		return nil, fmt.Errorf("internal server error")
	}

	url := workflowUrl + "/workflows/" + id
	res, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no workflows found")
	}
	defer res.Body.Close()

	text, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	body := models.Workflow{}
	err = json.Unmarshal(text, &body)
	if err != nil {
		return nil, err
	}

	return &body, nil
}

func DeleteWorkflow(id string) error {
	workflowUrl := os.Getenv("WORKFLOW_URL")
	if workflowUrl == "" {
		return fmt.Errorf("internal server error")
	}

	url := workflowUrl + "/workflows/" + id
	_, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return nil
}

func CreateWorkflow(input *models.WorkflowCreate) (*models.Workflow, error) {
	workflowUrl := os.Getenv("WORKFLOW_URL")
	if workflowUrl == "" {
		return nil, fmt.Errorf("internal server error")
	}

	byteArray, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	url := workflowUrl + "/workflows"
	res, err := http.NewRequest("POST", url, bytes.NewBuffer(byteArray))
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("workflow creation failed")
	}
	defer res.Body.Close()

	text, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	body := models.Workflow{}
	err = json.Unmarshal(text, &body)
	if err != nil {
		return nil, err
	}

	return &body, nil
}
