package service

import (
	openai_client "coderunner-service/internal/openai"
	"context"
	"encoding/base64"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

const formaterPrompt = "Your job is to format input data to fit data on a given example"
const translatorPrompt = "Your job is to translate input data to "
const generatePrompt = ""

func chatMessage(inputUser string, inputSystem string) (*string, error) {
	resp, err := openai_client.Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: inputSystem,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: inputUser,
				},
			},
		},
	)

	if err != nil {
		return nil, err
	}

	return &resp.Choices[0].Message.Content, nil
}

func LLMFormater(input string, example string) (*string, error) {
	return chatMessage(input, formaterPrompt+"\n Example:\n"+example)
}

func LLMTranslator(input string, languages []string) ([]*string, error) {
	var output []*string
	for _, language := range languages {
		translated, err := chatMessage(input, translatorPrompt+language)
		if err != nil {
			return nil, err
		}
		output = append(output, translated)
	}
	return output, nil
}

func LLMGenerate(input string) (*string, error) {
	return chatMessage(input, generatePrompt)
}

func LLMImage(input string) ([]byte, error) {
	request := openai.ImageRequest{
		Prompt:         input,
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
		N:              1,
	}

	response, err := openai_client.Client.CreateImage(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("Image creation error: ", err)

	}
	imgBytes, err := base64.StdEncoding.DecodeString(response.Data[0].B64JSON)
	if err != nil {
		return nil, fmt.Errorf("Base64 decode error: ", err)

	}
	return imgBytes, nil
}
