package middleware

import (
	"log/slog"

	"github.com/nyakovchuk/vsch_youtube_bot/internal/service"
	"github.com/nyakovchuk/vsch_youtube_bot/internal/shareddata"
	"gopkg.in/telebot.v4"
)

type BotManager interface {
	Logger() *slog.Logger
	TBot() *telebot.Bot
	LoggerInfo(telebot.Context)
	LoggerError(telebot.Context, error)
	Services() *service.Service
	SharedData() shareddata.Data
}
