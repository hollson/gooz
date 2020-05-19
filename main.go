// ------------------------------------------------------------------
// @ Author: hollson <hollson@live.cn>
// @ Date: 2019-12-01
// @ Version: 1.0.0
// ------------------------------------------------------------------

package main

import (
	"runtime"

	"github.com/hollson/deeplink/app"
	"github.com/hollson/deeplink/app/config"
	"github.com/hollson/deeplink/repo"
)

// 按顺序加载初始化项
func init() {
	config.Load()       // 加载配置
	repo.InitMysql()    // 初始化Mysql
	repo.InitPostgres() // 初始化Postgres
	repo.InitRedis()    // 初始化Redis
	config.InitLog()    // 初始化日志
}

func main() {
	runtime.Gosched()
	app.Run()
}
