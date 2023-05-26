package main

import (
	"telebot/core"
	"telebot/global"
	"telebot/initialize"

	"go.uber.org/zap"
)

func main() {
	// botToken := global.CONFIG.Bot.Token
	// bot, err := telebot.NewBot(botToken)
	// if err != nil {
	// 	log.Fatalf("Failed to create bot: %v", err)
	// }

	// go bot.InitBot()

	// c, err := canal.NewCanal(bot)
	// if err != nil {
	// 	log.Fatalf("Failed to create Canal: %v", err)
	// }

	// // Start canal
	// c.Run()

	// // Graceful shutdown
	// sigs := make(chan os.Signal, 1)
	// signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	// <-sigs

	// c.Close()
	// log.Println("Bye")

	global.VP = core.Viper()
	global.LOG = core.Zap()
	zap.ReplaceGlobals(global.LOG)
	global.DB = initialize.Gorm()
	initialize.Timer()
	if global.DB != nil {
		db, _ := global.DB.DB()
		defer db.Close()
	}

	initialize.NewRabbitMQPubSub("amqp://guest:guest@localhost:5672/", "videos")

	err := initialize.RunCanal(false)
	if err != nil {
		global.LOG.Error("canal run error: ", zap.Error(err))
	}
}
