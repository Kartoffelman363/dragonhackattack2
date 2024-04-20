package openai_client

import (
	"os"

	openai "github.com/sashabaranov/go-openai"
)

var Client *openai.Client

func Initialize() {
	Client = openai.NewClient(os.Getenv("OPENAI_API_KEY"))
}
