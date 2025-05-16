package telegram

import (
	"github.com/nyakovchuk/vsch_youtube_bot/internal/bot/telegram/handler"
)

func (b *Bot) Handlers() {
	b.bot.Handle("/start", handler.HandleStart(b))
	// b.bot.Handle("/getdb", handler.HandleGetDB(b), middleware.AdminOnly)
}
