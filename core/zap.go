package core

import (
	"fmt"
	"os"
	"telebot/core/internal"
	"telebot/global"
	"telebot/utils"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExits(global.CONFIG.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", global.CONFIG.Zap.Director)
		_ = os.Mkdir(global.CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
