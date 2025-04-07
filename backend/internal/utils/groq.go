package utils

import (
	"fmt"
	"os"

	"github.com/jpoz/groq"
)

type GroqClient struct {
	client *groq.Client
}

func NewGroqClient() *GroqClient {
	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		panic("GROQ_API_KEY not set")
	}

	client := groq.NewClient(groq.WithAPIKey(apiKey))
	return &GroqClient{client: client}
}

func (g *GroqClient) GetCompanySummary(company string, ticker string) (string, error) {
	prompt := fmt.Sprintf("Write a short one-paragraph summary about %s (%s). Format the output like: \"%s (%s) is a ...\"", company, ticker, company, ticker)

	resp, err := g.client.CreateChatCompletion(groq.CompletionCreateParams{
		Model: "llama3-8b-8192",
		Messages: []groq.Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	})
	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no choices returned from Groq API")
	}

	return resp.Choices[0].Message.Content, nil
}
