// -------------------------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.cn>
// @ Date: 2019-12-01
// @ Version: 1.0.0
// -------------------------------------------------------------------------------------

package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hollson/deeplink/app/config"
	"github.com/hollson/deeplink/app/monitor"
	_ "github.com/hollson/deeplink/repo"
)

var router *gin.Engine

func Init() {
	// 欢迎光临
	fmt.Printf(WELCOME, config.App.Version, config.App.Env, config.App.Name)
}

// 启动程序
func Run() {
	router = gin.Default()

	// 统计API和系统信息
	router.Use(monitor.ApiVisitHandler)
	router.GET("/help/stats", monitor.GetCurrentRunningStats)
	router.GET("/help/check", monitor.GetCurrentRunningStats) // 检查

	// s := &http.Server{

	RegisterRoute() // 注册路由表

	// 	Addr:           ":8080",
	// 	Handler:        router,
	// 	ReadTimeout:    5 * time.Second,
	// 	WriteTimeout:   5 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	//
	// s.ListenAndServe()  //开始监听

	if err := router.Run(config.App.Port); err != nil {
		panic(err)
	}
}
