package service

import (
	models "coderunner-service/pkg/models"
	"fmt"
	"reflect"
	"strconv"

	internalhandlers "coderunner-service/internal/internal_handlers"
	"coderunner-service/internal/stack"
)

func convert(data interface{}, dataType string) (interface{}, error) {
	switch dataType {
	case "string":
		return data.(string), nil
	case "int":
		val, err := strconv.ParseInt(data.(string), 10, 64)
		if err != nil {
			return nil, err
		}
		return int(val), nil
	case "float":
		val, err := strconv.ParseFloat(data.(string), 64)
		if err != nil {
			return nil, err
		}
		return int(val), nil
	case "hex":
		val, err := strconv.ParseInt(data.(string), 16, 64)
		if err != nil {
			return nil, err
		}
		return int(val), nil
	case "oct":
		val, err := strconv.ParseInt(data.(string), 8, 64)
		if err != nil {
			return nil, err
		}
		return int(val), nil
	default:
		return data, nil
	}
}

func StartParsing(workflow models.Workflow) error {
	globals := make(map[string]interface{})

	executeOrder := stack.New()
	executedBlocks := 0
	initial := len(workflow.Blocks.Workflows)
	for _, value := range workflow.Blocks.Workflows {
		executeOrder.Push(&value, executedBlocks)
	}

	for _, value := range workflow.InputVariables {
		data, err := convert(value.Value, value.Type)
		if err != nil {
			return err
		}
		globals[value.VarName] = data
	}

	for executeOrder.Len() > 0 {
		initial--
		current, counter := executeOrder.Pop()
		for _, value := range current.Blocks.Workflows {
			executeOrder.Push(&value, executedBlocks)
		}

		variables := make(map[string]interface{})

		for _, variable := range current.InputVariables {
			value, ok := globals[variable.Value]
			if !ok {
				if counter != executedBlocks || initial > 0 {
					executeOrder.Push(current, executedBlocks)
					continue
				} else {
					return fmt.Errorf("cannot find input variables for block_id: %s", current.ID)
				}
			}
			data, err := convert(value.(string), variable.Type)
			if err != nil {
				return err
			}
			variables[variable.VarName] = data
		}

		var output interface{}
		var err error
		switch current.Code {
		case "get_document":
			output, err = internalhandlers.GetDocumentByID(variables["_id"].(string))

		case "api_request":

			requestMethod, ok := variables["requestMethod"]
			if !ok {
				requestMethod = "POST"
			}
			contentType, ok := variables["contentType"]
			if !ok {
				contentType = "application/json"
			}
			output, err = ApiRequests(variables["url"].(string), variables["body"].(string), requestMethod.(string), contentType.(string))

		case "llm_formater":
			output, err = LLMFormater(variables["input"].(string), variables["example"].(string))

		case "llm_translator":
			output, err = LLMFormater(variables["input"].(string), variables["example"].(string))

		case "llm_generate_image":
			output, err = LLMImage(variables["input"].(string))

		case "llm_generate":
			output, err = LLMGenerate(variables["input"].(string))

		case "img_resize":
			output, err = ResizeImage(variables["inputBytes"].([]byte), variables["width"].(int), variables["height"].(int))

		default:
			continue
		}

		if err != nil {
			return err
		}

		if reflect.TypeOf(output).Kind() == reflect.Slice || reflect.TypeOf(output).Kind() == reflect.Array {
			outputValue := reflect.ValueOf(output)
			minLen := min(len(current.OutputVariables), outputValue.Len())
			for i := 0; i < minLen; i++ {
				globals[current.OutputVariables[i].VarName] = outputValue.Index(i).Interface()
			}
		} else {
			value := current.OutputVariables[0]
			globals[value.VarName] = output
		}

		executedBlocks++
	}

	return nil
}
