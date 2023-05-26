package initialize

import (
	"os"
	"telebot/global"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	return GormMysql()
}

func RegisterTables() {
	db := global.DB
	err := db.AutoMigrate()
	if err != nil {
		global.LOG.Error("register tables failed", zap.Error(err))
		os.Exit(0)
	}
	global.LOG.Info("register tables success")
}
