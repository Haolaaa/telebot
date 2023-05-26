package initialize

import (
	"telebot/global"
	"telebot/initialize/internal"

	"github.com/go-mysql-org/go-mysql/canal"
	"go.uber.org/zap"
)

func getCanalConfig() *canal.Config {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = global.CONFIG.Mysql.Path + ":" + global.CONFIG.Mysql.Port
	cfg.User = global.CONFIG.Mysql.Username
	cfg.Password = global.CONFIG.Mysql.Password
	cfg.Flavor = "mysql"
	cfg.Dump.TableDB = global.CONFIG.Mysql.DBname
	// cfg.Dump.Tables = []string{"hubserver.video_changes"}
	cfg.ParseTime = true
	return cfg
}

func setEventHandler(c *canal.Canal) {
	c.SetEventHandler(&internal.EventHandler{})
}

func RunCanal(isPos bool) error {
	cfg := getCanalConfig()
	c, err := canal.NewCanal(cfg)
	if err != nil {
		global.LOG.Error("canal configure error: ", zap.Error(err))
	}

	global.LOG.Info("canal is running")

	setEventHandler(c)

	if !isPos {
		return c.Run()
	}

	masterPos, err := c.GetMasterPos()
	if err != nil {
		global.LOG.Error("failed to get binlog position", zap.Error(err))
	}

	return c.RunFrom(masterPos)
}
