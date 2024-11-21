package dto

import tgData "github.com/telegram-mini-apps/init-data-golang"

type AuthResponse struct {
	Ok      bool        `json:"ok"`
	User    tgData.User `json:"user"`
	ChatID  string		`json:"chat_id"`
	Message string      `json:"message"`
}

type AuthRequest struct {
    InitData string `json:"initData"`
    IsMocked bool `json:"isMocked"`
}