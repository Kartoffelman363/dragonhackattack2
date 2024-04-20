package service

import (
	models "coderunner-service/pkg/models"

	"coderunner-service/internal/stack"
)

func StartParsing(workflow models.Workflow) ([]byte, error) {
	variables := make(map[string]models.Variable)

	executeOrder := stack.New()

	executeOrder.Push(&workflow)

	for executeOrder.Len() > 0 {
		current := executeOrder.Pop()
		for _, value := range current.InputVariables {
			variables[value.VarName] = value
		}
		for _, value := range current.OutputVariables {
			variables[value.VarName] = value
		}
		for _, value := range current.Blocks.Workflows {
			executeOrder.Push(&value)
		}

		switch current.Code {
		case "api_request":
			content, err := ApiRequests("", nil, "", "")
			if err != nil {
				return nil, err
			}
			_ = content
		case "llm_formater":
			content, err := LLMFormater("", "")
			if err != nil {
				return nil, err
			}
			_ = content
		case "llm_translator":
			content, err := LLMFormater("", "")
			if err != nil {
				return nil, err
			}
			_ = content
		case "llm_generate_image":
			content, err := LLMImage("")
			if err != nil {
				return nil, err
			}
			_ = content
		case "llm_generate":
			content, err := LLMGenerate("")
			if err != nil {
				return nil, err
			}
			_ = content
		case "img_resize":
			content, err := ResizeImage(nil, 0, 0)
			if err != nil {
				return nil, err
			}
			_ = content
		default:
			continue
		}
	}

	return nil, nil
}
