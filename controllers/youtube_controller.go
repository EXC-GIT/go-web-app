package controllers

import (
	"net/http"
	"sample-api/models"
	"sample-api/services"

	"github.com/gin-gonic/gin"
)

type YouTubeController struct {
	youtubeService *services.YouTubeService
}

func NewYouTubeController(youtubeService *services.YouTubeService) *YouTubeController {
	return &YouTubeController{
		youtubeService: youtubeService,
	}
}

func (yc *YouTubeController) ExtractAudio(c *gin.Context) {
	var req models.ExtractAudioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ExtractAudioResponse{
			Success: false,
			Message: "Invalid request: " + err.Error(),
		})
		return
	}

	filePath, err := yc.youtubeService.ExtractAudio(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ExtractAudioResponse{
			Success: false,
			Message: "Failed to extract audio: " + err.Error(),
		})
		return
	}

	// Return file download
	if err := yc.youtubeService.ServeAudioFile(c, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, models.ExtractAudioResponse{
			Success: false,
			Message: "Failed to serve file: " + err.Error(),
		})
		return
	}
}
