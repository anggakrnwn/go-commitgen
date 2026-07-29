package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/anggakrnwn/go-commitgen/internal/domain"
	"google.golang.org/genai"
)

type GeminiClient struct {
	client *genai.Client
	model  string
}

func NewGeminiClient(apiKey, defaultModel string) (*GeminiClient, error) {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to  create gemini client: %w", err)
	}

	if defaultModel == "" {
		defaultModel = os.Getenv("DEFAULTMODEL")
	}

	return &GeminiClient{
		client: client,
		model:  defaultModel,
	}, nil
}

func (g *GeminiClient) GenerateCommitMessage(req domain.LLMRequest) (*domain.LLMResponses, error) {
	ctx := context.Background()

	fullText := fmt.Sprintf("%s\n\nGit Diff:\n%s", req.Prompt, req.DiffText)

	config := &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
		ResponseSchema: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"CommitType": {
					Type:        genai.TypeString,
					Description: "Tipe commit konvensional seperti feat, fix, refactor, chore, docs, dll.",
				},
				"CommitScope": {
					Type:        genai.TypeString,
					Description: "Cakupan/modul kode yang diubah, contoh: repository, usecase, auth.",
				},
				"CommitDescription": {
					Type:        genai.TypeString,
					Description: "Penjelasan singkat mengenai perubahan yang terjadi.",
				},
				"CommitBody": {
					Type:        genai.TypeString,
					Description: "Penjelasan detail atau alasan perubahan jika diperlukan (opsional).",
				},
			},
			Required: []string{"CommitType", "CommitDescription"},
		},
	}

	response, err := g.client.Models.GenerateContent(ctx, g.model, genai.Text(fullText), config)
	if err != nil {
		return nil, fmt.Errorf("failed to generate content: %w", err)
	}

	jsontext := response.Text()
	if jsontext == "" {
		return nil, fmt.Errorf("empty response")
	}

	var result domain.LLMResponses

	if err := json.Unmarshal([]byte(jsontext), &result); err != nil {
		return nil, fmt.Errorf("failed to parse json response: %w, raw text: %s", err, jsontext)
	}

	return &result, nil

}
