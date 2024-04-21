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

type varMap struct {
	Type  string      `json:"type" bson:"type"`
	Value interface{} `json:"value" bson:"value"`
}

func StartParsing(workflow models.Workflow) ([]models.Output, error) {

	globals := make(map[string]varMap)

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
		globals[value.Id] = varMap{Type: value.Type, Value: data}
	}

	var returnedOutput []models.Output

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
				globals[value.Id] = varMap{Type: value.Type, Value: data}
			}
		} else {
			fmt.Println(executeOrder.String())
			variables := make(map[string]varMap)
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
			_type := "string"
			switch current.Code {
			case "get_document":
				output, err = internalhandlers.GetDocumentByID(variables["id"].Value.(string))

			case "api_request":

				pointer, err = ApiRequests(variables["url"].Value.(string), variables["body"].Value.(string), "POST", "application/json")
				output = *pointer
			case "llm_formater":
				pointer, err = LLMFormater(variables["input"].Value.(string), variables["example"].Value.(string))
				output = *pointer
			case "llm_translator":
				pointer, err = LLMFormater(variables["input"].Value.(string), variables["language"].Value.(string))
				output = *pointer
			case "llm_generate_image":
				pointer, err = LLMImage(variables["input"].Value.(string))
				output = *pointer
				_type = "image"
			case "llm_generate_image_prompt":
				pointer, err = LLMImagePrompt(variables["input"].Value.(string))
				output = *pointer
			case "llm_generate_keyword":
				pointer, err = LLMKeywords(variables["input"].Value.(string))
				output = *pointer
			case "llm_generate":
				pointer, err = LLMGenerate(variables["input"].Value.(string))
				output = *pointer
			case "img_resize":
				output, err = ResizeImage(variables["inputbytes"].Value.([]byte), variables["width"].Value.(int), variables["height"].Value.(int))

			case "output":
				for key, value := range variables {
					fmt.Println(key, value)
					returnedOutput = append(returnedOutput, models.Output{Type: value.Type, Value: value.Value.(string)})

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
					globals[current.OutputVariables[i].VarName] = varMap{Type: _type, Value: outputValue.Index(i).Interface()}
				}
			} else {
				value := current.OutputVariables[0]
				globals[value.VarName] = varMap{Type: _type, Value: output}
			}
		}
		executedBlocks++
	}

	jsonString, _ := json.Marshal(returnedOutput)
	fmt.Print("OUTPUT ")
	fmt.Println(string(jsonString))

	return returnedOutput, nil
}
