package initialize

import (
	"telebot/global"
	"telebot/initialize/internal"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TelegramBot() (*internal.Bot, error) {
	var api *tgbotapi.BotAPI
	var err error

	api, err = tgbotapi.NewBotAPI(global.CONFIG.Bot.Token)
	if err != nil {
		return nil, err
	}

	bot := &internal.Bot{}

	bot.SetBotAPI(api)

	return &internal.Bot{
		Username: api.Self.UserName,
	}, nil
}
