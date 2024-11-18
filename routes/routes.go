package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/paigexx/telegram-go-server/routes/auth"
	"github.com/paigexx/telegram-go-server/routes/files"
)

func ApplyRoutes(r *gin.Engine) {
    api := r.Group("/")
    auth.NewHandler(api)
    files.NewHandler(api)
}