package main

import (
	"runtime"

	"github.com/hollson/deeplink/app"
	"github.com/hollson/deeplink/app/config"
	"github.com/hollson/deeplink/app/logger"
	"github.com/hollson/deeplink/repo"
)

// 按顺序加载初始化项
func init() {
	config.Init()
	app.Init()
	repo.Init()
	logger.Init()
}

//go:generate go build -o deeplink
func main() {
	runtime.Gosched()
	app.Run()
}
