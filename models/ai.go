package models

// AIPromptRequest represents a request to send a prompt to the AI platform
type AIPromptRequest struct {
	Prompt   string `json:"prompt" binding:"required"`
	Provider string `json:"provider,omitempty"` // openai, google, anthropic
}

// AIPromptResponse represents the response from the AI platform
type AIPromptResponse struct {
	Success  bool   `json:"success"`
	Message  string `json:"message,omitempty"`
	Response string `json:"response,omitempty"`
	Error    string `json:"error,omitempty"`
}

// AIAnalysisRequest represents a request to analyze YouTube content
type AIAnalysisRequest struct {
	Content      string `json:"content" binding:"required"`
	Provider     string `json:"provider,omitempty"`      // openai, google, anthropic
	AnalysisType string `json:"analysis_type,omitempty"` // summary, sentiment, keywords, etc.
}

// AIAnalysisResponse represents the response from AI analysis
type AIAnalysisResponse struct {
	Success  bool   `json:"success"`
	Message  string `json:"message,omitempty"`
	Analysis string `json:"analysis,omitempty"`
	Error    string `json:"error,omitempty"`
}
