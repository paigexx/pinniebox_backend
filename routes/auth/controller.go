package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paigexx/telegram-go-server/dto"
	"github.com/paigexx/telegram-go-server/services"
)

type Handler struct {
	service services.AuthService
}

func newHandler(authService services.AuthService) *Handler {
	return &Handler{
		service: authService,
	}
}


func (h Handler) Authenticate(c *gin.Context) {
    input := dto.AuthRequest{}

    if err := c.ShouldBind(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
        return
    }

    result, err := h.service.Authenticate(c, &input.InitData, input.IsMocked)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Authentication failed"})
        return
    }
    c.JSON(http.StatusOK, result)
}