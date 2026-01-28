package services

import (
	"fmt"
	"log"
	"os"

	"sample-api/services/providers"
)

// AIService handles communication with AI platforms
type AIService struct {
	provider  providers.AIProvider
	apiKey    string
	modelName string
}

// NewAIService creates a new AI service with the specified provider
func NewAIService(providerType string) *AIService {
	apiKey := os.Getenv("AI_API_KEY")
	googleModel := os.Getenv("GOOGLE_MODEL")
	if apiKey == "" {
		log.Println("Warning: AI_API_KEY environment variable not set")
	}

	var provider providers.AIProvider

	switch providerType {
	case "openai":
		provider = &providers.OpenAIProvider{
			APIKey:    apiKey,
			ModelName: "gpt-3.5-turbo",
		}
	case "google":
		provider = &providers.GoogleAIProvider{
			APIKey:    apiKey,
			ModelName: googleModel,
		}
	case "anthropic":
		provider = &providers.AnthropicProvider{
			APIKey:    apiKey,
			ModelName: "claude-3-sonnet-20240229",
		}
	default:
		log.Printf("Unknown provider: %s, defaulting to OpenAI", providerType)
		provider = &providers.OpenAIProvider{
			APIKey:    apiKey,
			ModelName: "gpt-3.5-turbo",
		}
	}

	return &AIService{
		provider:  provider,
		apiKey:    apiKey,
		modelName: providerType,
	}
}

// PromptAI sends a prompt to the AI platform and returns the response
func (as *AIService) PromptAI(prompt string) (string, error) {
	if as.apiKey == "" {
		return "", fmt.Errorf("AI API key not configured")
	}

	return as.provider.PromptAI(prompt)
}

// AnalyzeYouTubeContent uses AI to analyze YouTube audio/content
func (as *AIService) AnalyzeYouTubeContent(content string) (string, error) {
	prompt := fmt.Sprintf("Analyze the following YouTube content and provide a summary:\n\n%s", content)
	return as.PromptAI(prompt)
}

// TranscribeAudio uses AI to transcribe audio content
func (as *AIService) TranscribeAudio(audioPath string) (string, error) {
	prompt := fmt.Sprintf("Transcribe the audio file at: %s", audioPath)
	return as.PromptAI(prompt)
}

// GenerateSummary generates a summary of provided text using AI
func (as *AIService) GenerateSummary(text string, length string) (string, error) {
	prompt := fmt.Sprintf("Generate a %s summary of the following text:\n\n%s", length, text)
	return as.PromptAI(prompt)
}
