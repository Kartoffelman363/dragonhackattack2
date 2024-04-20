package modules

// when changing file make sure to update across repo

type Status struct {
	StatusCode string `json:"statuscode"`
}

type DocumentDeletionRequest struct {
	ID string `json:"_id"`
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

type Variable struct {
	VarName string `json:"varname"`
	Type    string `json:"type"`
	Value   string `json:"value"`
	Example string `json:"example"`
}

type MetaData struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Workflow struct {
	ID              string     `json:"id"`
	Name            string     `json:"name"`
	Metadata        MetaData   `json:"metadata"`
	InputVariables  []Variable `json:"input_variables"`
	OutputVariables []Variable `json:"output_variables"`
	Code            string     `json:"code"`
	Blocks          Workflows  `json:"blocks"`
}

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
