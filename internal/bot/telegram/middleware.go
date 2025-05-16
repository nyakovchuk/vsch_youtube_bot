package telegram

import "github.com/nyakovchuk/vsch_youtube_bot/internal/bot/telegram/middleware"

func (b *Bot) Middleware() {
	middleware.CheckUser(b)
}
