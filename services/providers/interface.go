package providers

// AIProvider defines the interface for AI platform providers
type AIProvider interface {
	PromptAI(prompt string) (string, error)
}
