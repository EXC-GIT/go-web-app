package services

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type YouTubeService struct {
	tempDir string
}

func NewYouTubeService() *YouTubeService {
	tempDir := filepath.Join(os.TempDir(), "youtube_audio")
	os.MkdirAll(tempDir, 0755)
	return &YouTubeService{
		tempDir: tempDir,
	}
}

func (ys *YouTubeService) ExtractAudio(url string) (string, error) {
	// Generate unique filename
	fileID := uuid.New().String()
	outputPath := filepath.Join(ys.tempDir, fmt.Sprintf("%s.mp3", fileID))

	// yt-dlp command to extract audio as MP3
	cmd := exec.Command("yt-dlp",
		"--extract-audio",
		"--audio-format", "mp3",
		"--audio-quality", "192K",
		"--output", outputPath,
		"--no-playlist",
		url,
	)

	// Execute command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("yt-dlp failed: %v, output: %s", err, string(output))
	}

	// Verify file was created
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		return "", fmt.Errorf("audio file was not created")
	}

	return outputPath, nil
}

func (ys *YouTubeService) ServeAudioFile(c *gin.Context, filePath string) error {
	// Get filename from path
	filename := filepath.Base(filePath)
	if !strings.HasSuffix(filename, ".mp3") {
		filename += ".mp3"
	}

	// Set headers for file download
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	c.Header("Content-Type", "audio/mpeg")
	c.Header("Content-Transfer-Encoding", "binary")

	// Serve the file
	c.File(filePath)

	// Clean up file after serving (with delay to ensure download completes)
	go func() {
		time.Sleep(10 * time.Second)
		os.Remove(filePath)
	}()

	return nil
}

func (ys *YouTubeService) Cleanup() {
	os.RemoveAll(ys.tempDir)
}
