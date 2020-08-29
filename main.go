package main

import (
	"runtime"

	"github.com/hollson/gooz/app"
	"github.com/hollson/gooz/app/config"
	"github.com/hollson/gooz/app/logger"
	"github.com/hollson/gooz/repo"
)

// 按顺序加载初始化项
func init() {
	config.Init()
	app.Init()
	repo.Init()
	logger.Init()
}

//go:generate go build -o gooz
func main() {
	runtime.Gosched()
	app.Run()
}
