package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/paigexx/telegram-go-server/routes"
)


func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Println("Error loading .env file")
    }

    // Ensure the Telegram bot token is set
    telegramBotToken := os.Getenv("TELEGRAM_BOT_TOKEN")
    if telegramBotToken == "" {
        log.Fatal("TELEGRAM_BOT_TOKEN is not set")
    }


	// Configure CORS settings
	corsConfig := cors.Config{
		AllowOrigins:     []string{os.Getenv("WEB_APP_URL")}, 
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Requested-With", "Accept", "Origin"},
		AllowCredentials: true,
	}

	// Create a Gin router with default middleware
	r := gin.Default()
	r.Use(cors.New(corsConfig))
	r.RedirectTrailingSlash = false
	routes.ApplyRoutes(r)

    // Start the server
    port := ":8080"
    log.Printf("Server is running on port %s\n", port)
    if err := r.Run(port); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
