package providers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Google AI API request/response structures
type GoogleAIPart struct {
	Text string `json:"text"`
}

type GoogleAIContent struct {
	Parts []GoogleAIPart `json:"parts"`
}

type GoogleAIGenerationConfig struct {
	Temperature     float64 `json:"temperature,omitempty"`
	MaxOutputTokens int     `json:"maxOutputTokens,omitempty"`
}

type GoogleAIRequest struct {
	Contents         []GoogleAIContent        `json:"contents"`
	GenerationConfig GoogleAIGenerationConfig `json:"generationConfig,omitempty"`
}

type GoogleAICandidate struct {
	Content GoogleAIContent `json:"content"`
}

type GoogleAIResponse struct {
	Candidates []GoogleAICandidate `json:"candidates"`
	Error      struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// GoogleAIProvider implements AIProvider for Google AI
type GoogleAIProvider struct {
	APIKey    string
	ModelName string
}

// PromptAI sends a prompt to Google AI API
func (gp *GoogleAIProvider) PromptAI(prompt string) (string, error) {
	if prompt == "" {
		return "", fmt.Errorf("prompt cannot be empty")
	}

	if gp.APIKey == "" {
		return "", fmt.Errorf("Google AI API key not set")
	}

	// Create request payload
	reqPayload := GoogleAIRequest{
		Contents: []GoogleAIContent{
			{
				Parts: []GoogleAIPart{
					{
						Text: prompt,
					},
				},
			},
		},
	}

	reqPayload.GenerationConfig.Temperature = 0.7
	reqPayload.GenerationConfig.MaxOutputTokens = 1000

	// Convert to JSON
	jsonData, err := json.Marshal(reqPayload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	// Make HTTP request
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1/models/%s:generateContent?key=%s", gp.ModelName, gp.APIKey)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to call Google AI API: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	// Parse response
	var googleResp GoogleAIResponse
	err = json.Unmarshal(body, &googleResp)
	if err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	// Check for errors
	if googleResp.Error.Message != "" {
		return "", fmt.Errorf("Google AI error: %s", googleResp.Error.Message)
	}

	if len(googleResp.Candidates) == 0 {
		return "", fmt.Errorf("no response from Google AI")
	}

	if len(googleResp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no content in Google AI response")
	}

	response := googleResp.Candidates[0].Content.Parts[0].Text
	log.Printf("Google AI API called with model: %s", gp.ModelName)

	return response, nil
}
