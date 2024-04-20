package modules

// when changing file make sure to update across repo

type Status struct {
	StatusCode string `json:"statuscode"`
}

type DocumentDeletionRequest struct {
	ID string `json:"id"`
}

type DocumentGetByIDRequest struct {
	ID string `json:"id"`
}

type WorkflowDeletionRequest struct {
	ID string `json:"id"`
}

type WorkflowGetByIDRequest struct {
	ID string `json:"id"`
}

// TODO: Expand to include workflow
type Workflow struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// TODO: Expand to include document
type Document struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Workflows struct {
	Workflows []Workflow `json:"workflows"`
}

type Documents struct {
	Documents []Document `json:"documents"`
}
