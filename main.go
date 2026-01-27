package main

import (
	"log"
	"net"
	"os"
	"strconv"

	"sample-api/controllers"
	"sample-api/models"
	"sample-api/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize database
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate models
	db.AutoMigrate(&models.User{})

	// Initialize services
	userService := services.NewUserService(db)
	youtubeService := services.NewYouTubeService()
	userController := controllers.NewUserController(userService)
	youtubeController := controllers.NewYouTubeController(youtubeService)

	// Setup Gin router
	r := gin.Default()

	// CORS middleware
	r.Use(cors.Default())

	// Routes
	r.GET("/users", userController.GetUsers)
	r.POST("/users", userController.CreateUser)
	r.POST("/extract-audio", youtubeController.ExtractAudio)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Try to bind to the port, if occupied, find a random available port
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		// Port is occupied, find a random available port
		listener, err = net.Listen("tcp", ":0")
		if err != nil {
			log.Fatal("Failed to find available port:", err)
		}
		port = strconv.Itoa(listener.Addr().(*net.TCPAddr).Port)
	}
	listener.Close()

	log.Printf("Server starting on port %s", port)
	r.Run(":" + port)
}
