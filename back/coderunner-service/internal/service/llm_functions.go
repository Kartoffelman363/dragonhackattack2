package service

import (
	openai_client "coderunner-service/internal/openai"
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

const imagePrompt = "Your job is to create image prompt based on input data "
const keywordsPrompt = "Your job is to extract keywords from input data "
const summarizePrompt = "Your job is to summarize input data"
const formaterPrompt = "Your job is to format input data to fit data on a given example"
const translatorPrompt = "Your job is to translate input data to "
const generatePrompt = ""

func ChatMessage(inputUser string, inputSystem string) (*string, error) {
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

func LLMImagePrompt(input string) (*string, error) {
	return ChatMessage(input, imagePrompt)
}

func LLMKeywords(input string) (*string, error) {
	return ChatMessage(input, keywordsPrompt)
}

func LLMSumarize(input string) (*string, error) {
	return ChatMessage(input, summarizePrompt)
}

func LLMFormater(input string, example string) (*string, error) {
	return ChatMessage(input, formaterPrompt+"\n Example:\n"+example)
}

func LLMTranslator(input string, language string) (*string, error) {
	translated, err := ChatMessage(input, translatorPrompt+language)
	if err != nil {
		return nil, err
	}
	return translated, nil
}

func LLMGenerate(input string) (*string, error) {
	return ChatMessage(input, generatePrompt)
}

func LLMImage(input string) (*string, error) {
	request := openai.ImageRequest{
		Prompt:         input,
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
		N:              1,
	}

	response, err := openai_client.Client.CreateImage(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("image creation error: %v", err)
	}

	return &response.Data[0].B64JSON, nil
}
