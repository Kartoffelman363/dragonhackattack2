package service

import (
	models "coderunner-service/pkg/models"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	internalhandlers "coderunner-service/internal/internal_handlers"
	"coderunner-service/internal/queue"
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

func StartParsing(workflow models.Workflow) ([]byte, error) {
	globals := make(map[string]interface{})

	executeOrder := queue.New()
	executedBlocks := 0
	initial := len(workflow.Blocks)
	val, _ := json.Marshal(workflow.Blocks)
	fmt.Println(string(val))
	for _, value := range workflow.Blocks {
		fmt.Println(value)
		temp := value
		executeOrder.Enqueue(&temp, executedBlocks)
	}
	fmt.Println(executeOrder.String())

	for _, value := range workflow.InitialVariables {
		data, err := convert(value.Value, value.Type)
		if err != nil {
			return nil, err
		}
		globals[value.Id] = data
	}

	returnedOutput := make(map[string]models.Output)

	for executeOrder.Len() > 0 {
		jsonString, _ := json.Marshal(globals)
		fmt.Print("GLOBALS ")
		fmt.Println(string(jsonString))
		initial--
		current, counter := executeOrder.Dequeue()
		fmt.Println("eXECUTING BLOCK TYPE", current.Code)
		if current.Code == "constants" {
			for _, value := range current.OutputVariables {
				data, err := convert(value.Value, value.Type)
				if err != nil {
					return nil, err
				}
				globals[value.Id] = data
			}
		} else {
			fmt.Println(executeOrder.String())
			variables := make(map[string]interface{})
			breakage := false
			for _, variable := range current.InputVariables {
				fmt.Println(variable.VarName, " ", variable.Value, " ", variable.Id)
				value, ok := globals[variable.Value]
				if !ok {
					if counter != executedBlocks || initial > 0 {
						fmt.Println("enqued")
						executeOrder.Enqueue(current, executedBlocks)
						breakage = true
						break
					} else {
						return nil, fmt.Errorf("cannot find input variables for block_id: %s", current.ID)
					}
				}
				fmt.Println("OK")
				/*data, err := convert(value.(string), variable.Type)
				if err != nil {
					return nil, err
				}*/
				fmt.Println("OK")
				variables[strings.ToLower(variable.VarName)] = value
			}
			if breakage {
				continue
			}
			jsonString, _ := json.Marshal(variables)
			fmt.Print("VARIABLES ")
			fmt.Println(string(jsonString))
			var pointer *string
			var output interface{}
			var err error
			switch current.Code {
			case "get_document":
				output, err = internalhandlers.GetDocumentByID(variables["id"].(string))

			case "api_request":

				requestMethod, ok := variables["requestmethod"]
				if !ok {
					requestMethod = "POST"
				}
				contentType, ok := variables["contenttype"]
				if !ok {
					contentType = "application/json"
				}
				output, err = ApiRequests(variables["url"].(string), variables["body"].(string), requestMethod.(string), contentType.(string))
			case "llm_formater":
				pointer, err = LLMFormater(variables["input"].(string), variables["example"].(string))
				output = *pointer
			case "llm_translator":
				pointer, err = LLMFormater(variables["input"].(string), variables["language"].(string))
				output = *pointer
			case "llm_generate_image":
				output, err = LLMImage(variables["input"].(string))

			case "llm_generate_image_prompt":
				pointer, err = LLMImagePrompt(variables["input"].(string))
				output = *pointer
			case "llm_generate_keyword":
				pointer, err = LLMKeywords(variables["input"].(string))
				output = *pointer
			case "llm_generate":
				pointer, err = LLMGenerate(variables["input"].(string))
				output = *pointer
			case "img_resize":
				output, err = ResizeImage(variables["inputbytes"].([]byte), variables["width"].(int), variables["height"].(int))

			case "output":
				for key, value := range variables {
					fmt.Println(key, value)
					byteValue, ok := value.([]byte)
					if !ok {
						returnedOutput[key] = models.Output{Type: "string", Value: value.(string)}
					} else {
						returnedOutput[key] = models.Output{Type: "image", Value: string(byteValue)}
					}
				}
				continue
			default:
				continue
			}

			if err != nil {
				return nil, err
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
		}
		executedBlocks++
	}

	jsonString, _ := json.Marshal(returnedOutput)
	fmt.Print("OUTPUT ")
	fmt.Println(string(jsonString))

	return jsonString, nil
}
