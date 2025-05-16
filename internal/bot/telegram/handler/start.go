package handler

import (
	"gopkg.in/telebot.v4"
)

func HandleStart(bm BotManager) func(c telebot.Context) error {
	return func(c telebot.Context) error {
		bm.LoggerInfo(c)

		return c.Send("Hello", &telebot.SendOptions{
			ParseMode:             telebot.ModeHTML,
			DisableWebPagePreview: true,
		})
	}
}
