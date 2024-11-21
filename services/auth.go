package services

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/paigexx/telegram-go-server/dto"
	tgData "github.com/telegram-mini-apps/init-data-golang"
)

type AuthService struct{}

func NewAuthService() *AuthService {
    return &AuthService{}
}


func (s *AuthService) Authenticate(c *gin.Context, initData *string, isMocked bool) (dto.AuthOutput, error) {
	if initData == nil && !isMocked {
        return dto.AuthOutput{}, errors.New("initData is required")
    }

    // Get the Telegram Bot Token from environment variables
    telegramBotToken := os.Getenv("TELEGRAM_BOT_TOKEN")
    if telegramBotToken == "" {
		return dto.AuthOutput{}, errors.New("telegram bot token is not set")

    }

    // Handle mocked data for testing
    if isMocked  {
        mockUserData := tgData.User{
            ID:         123456789,
            FirstName: 	"Test",
            LastName:  "User",
            Username:   "testuser",
            PhotoURL:  "https://www.gravatar.com/avatar",
        }

        response := dto.AuthOutput{
            User:    mockUserData,
			ChatID:  "123456789",
            Message: "Using mocked data",
        }
        return response, nil
    }

    // Define expiration time for initData (e.g., 24 hours)
    expiration := 24 * time.Hour


	if initData != nil {
		// Validate the initData with the Telegram Bot Token and expiration time
		err := tgData.Validate(*initData, telegramBotToken, expiration)
		if err != nil {
			log.Println("Error validating initData:", err)
			return dto.AuthOutput{}, errors.New("invalid initData")
		}

		// Parse the initData to get user data
		initDataParsed, err := tgData.Parse(*initData)
		if err != nil {
			log.Println("Error parsing initData:", err)
			return dto.AuthOutput{}, errors.New("failed to parse initData")
		}
		// Respond with the parsed initData
		response := dto.AuthOutput{
			User:     initDataParsed.User,
			ChatID:   fmt.Sprint(initDataParsed.ChatInstance),
			Message:  "Using parsed data",
		}
		return response, nil
	}
	return dto.AuthOutput{}, errors.New("invalid initData")
}