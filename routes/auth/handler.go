package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/paigexx/telegram-go-server/services"
)


func NewHandler(r *gin.RouterGroup) {
	auth := r.Group("/auth")

	service := services.NewAuthService()
	h := newHandler(*service)

	auth.POST("", h.Authenticate)
	
}

