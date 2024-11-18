package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/paigexx/telegram-go-server/dto"
)

type FilesService struct{}

func NewFilesService() *FilesService {
    return &FilesService{}
}


func (s *FilesService) Upload(c *gin.Context, file multipart.File, fileName string, telegramID string, chatID string, chatType string) (error) {
    // Create a buffer to hold the multipart form data for Pinata
    var buf bytes.Buffer
    writer := multipart.NewWriter(&buf)

    // Create a form file field named "file"
    part, err := writer.CreateFormFile("file", fileName)
    if err != nil {
        return fmt.Errorf("error creating form file: %s", err)
    }

    // Copy the uploaded file data to the form file field
    _, err = io.Copy(part, file)
    if err != nil {
        return fmt.Errorf("error copying file data: %s", err)
    }

	// Create a map with your key-value pairs
	keyvaluesData := map[string]interface{}{
    "chat_id":     chatID,
    "telegram_id": telegramID,
    "chat_type":   chatType,
}

	// Marshal the map into a JSON string
	keyvaluesJSON, err := json.Marshal(keyvaluesData)
	if err != nil {
		return fmt.Errorf("error marshaling keyvalues: %s", err)
	}

	// Write the JSON string to the form field
	err = writer.WriteField("keyvalues", string(keyvaluesJSON))
	if err != nil {
		return fmt.Errorf("error writing keyvalues field: %s", err)
	}

    // Close the writer to finalize the multipart form data
    err = writer.Close()
    if err != nil {
        return fmt.Errorf("error closing writer: %s", err)
    }

    // Continue with the rest of your code...
    // Create a new POST request to Pinata's file upload endpoint
    url := "https://uploads.devpinata.cloud/v3/files"
    req, err := http.NewRequest("POST", url, &buf)
    if err != nil {
        return fmt.Errorf("error creating request: %s", err)
    }

    // Set the appropriate headers, including your Pinata JWT token
    jwt := os.Getenv("PINATA_JWT")
    req.Header.Set("Content-Type", writer.FormDataContentType())
    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwt)) // Replace with your actual token

    // Send the request to Pinata
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("error sending request: %s", err)
    }
    defer resp.Body.Close()

	// Read the response from Pinata
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error uploading file: %s", responseBody)
	}

	var pinataResp dto.FileUploadResponse
	err = json.Unmarshal(responseBody, &pinataResp)
	if err != nil {
		return fmt.Errorf("error unmarshaling response: %s", err)
	}
	return nil
}

func (s *FilesService) List(c *gin.Context, telegramID string, chatID string, pageToken string) (dto.ListFilesResponse, error) {

	url := fmt.Sprintf(`https://api.devpinata.cloud/v3/files?pageToken=%v&metadata[telegram_id]=%v&metadata[chat_id]=%v`, pageToken, telegramID, chatID)
	fmt.Printf("url: %v", url)
	req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return dto.ListFilesResponse{}, fmt.Errorf("error creating request: %s", err)
    }

    jwt := os.Getenv("PINATA_JWT")
    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwt)) // Replace with your actual token

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return dto.ListFilesResponse{}, fmt.Errorf("error sending request: %s", err)
    }
    defer resp.Body.Close()

    responseBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return dto.ListFilesResponse{}, fmt.Errorf("error reading response: %s", err)
    }
    if resp.StatusCode != http.StatusOK {
        return dto.ListFilesResponse{}, fmt.Errorf("error listing files: %s", responseBody)
    }

    var pinataResp dto.ListFilesResponse
    err = json.Unmarshal(responseBody, &pinataResp)
    if err != nil {
        return dto.ListFilesResponse{}, fmt.Errorf("error unmarshaling response: %s", err)
    }

    return pinataResp, nil
}

