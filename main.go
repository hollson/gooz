package main

import (
	"runtime"
	"time"

	"github.com/hollson/deeplink/app"
	"github.com/hollson/deeplink/app/config"
	"github.com/hollson/deeplink/app/logger"
	"github.com/hollson/deeplink/repo"
	"github.com/sirupsen/logrus"
)

// 按顺序加载初始化项
func init() {
	config.Init() // 初始化配置
	app.Init()    // 初始化应用
	repo.Init()   // 初始化数据
	logger.Init() // 初始化日志
}

func main() {
	runtime.Gosched()
	go func() {
		for {
			time.Sleep(time.Second)
			logrus.Errorln("<br/>ddddd")
			logrus.Infoln("eeeeee")
		}

	}()
	app.Run()
}
