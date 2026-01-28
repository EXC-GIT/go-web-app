package controllers

import (
	"net/http"

	"sample-api/models"
	"sample-api/services"

	"github.com/gin-gonic/gin"
)

type AIController struct {
	aiService *services.AIService
}

// NewAIController creates a new AI controller
func NewAIController(aiService *services.AIService) *AIController {
	return &AIController{
		aiService: aiService,
	}
}

// PromptAI handles requests to prompt the AI platform
func (ac *AIController) PromptAI(c *gin.Context) {
	var req models.AIPromptRequest

	// Bind JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.AIPromptResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	// Call AI service
	response, err := ac.aiService.PromptAI(req.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.AIPromptResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.AIPromptResponse{
		Success:  true,
		Message:  "Prompt processed successfully",
		Response: response,
	})
}

// AnalyzeYouTubeContent handles requests to analyze YouTube content using AI
func (ac *AIController) AnalyzeYouTubeContent(c *gin.Context) {
	var req models.AIAnalysisRequest

	// Bind JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.AIAnalysisResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	// Call AI service
	response, err := ac.aiService.AnalyzeYouTubeContent(req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.AIAnalysisResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.AIAnalysisResponse{
		Success:  true,
		Message:  "Content analyzed successfully",
		Analysis: response,
	})
}

// GenerateSummary handles requests to generate a summary using AI
func (ac *AIController) GenerateSummary(c *gin.Context) {
	var req models.AIAnalysisRequest

	// Bind JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.AIAnalysisResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	// Get summary length from query param (short, medium, long)
	length := c.DefaultQuery("length", "medium")

	// Call AI service
	response, err := ac.aiService.GenerateSummary(req.Content, length)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.AIAnalysisResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.AIAnalysisResponse{
		Success:  true,
		Message:  "Summary generated successfully",
		Analysis: response,
	})
}
