package initialize

import (
	_ "telebot/config"
	"telebot/global"

	_ "github.com/robfig/cron/v3"
)

func Timer() {
	if global.CONFIG.Timer.Start {
		// for i := range global.CONFIG.Timer.Detail {
		// go func(detail config.Detail) {
		// 	var option []cron.Option
		// 	if global.CONFIG.Timer.WithSeconds {
		// 		option = append(option, cron.WithSeconds())
		// 	}

		// 	_, err := global.Timer.AddTaskByFunc(detail.Name, detail.Spec, detail.Func, option...)
		// 	if err != nil {
		// 		fmt.Println("定时任务添加失败", err)
		// 	}
		// }(global.CONFIG.Timer.Detail[i])
		// }
	}
}
