package files

import (
	"github.com/gin-gonic/gin"
	"github.com/paigexx/telegram-go-server/services"
)


func NewHandler(r *gin.RouterGroup) {
	files := r.Group("/files")

	service := services.NewFilesService()
	h := newHandler(*service)

	files.POST("", h.Upload)
	files.GET("/:telegram_id/:chat_id", h.List)
	files.GET("signedUrl/:cid", h.GetSignedUrl)
}

