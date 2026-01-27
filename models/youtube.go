package models

type ExtractAudioRequest struct {
	URL string `json:"url" binding:"required"`
}

type ExtractAudioResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	FileURL string `json:"file_url,omitempty"`
}
