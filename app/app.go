// -------------------------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.cn>
// @ Date: 2019-12-01
// @ Version: 1.0.0
// -------------------------------------------------------------------------------------

package app

import (
	"github.com/gin-gonic/gin"
	"github.com/hollson/deeplink/app/config"
	_ "github.com/hollson/deeplink/repo"
)

// // http://www.network-science.de/ascii/
// const Welcome = `
// Deeplink is a internal applications <XX科技.保留版权>
//        _              _ _      _
//     __| |___ ___ _ __| (_)_ _ | |__
//    / _  / -_) -_) '_ \ | | ' \| / /
//    \____\___\___| .__/_|_|_||_|_\_\   %s(%s)
//                 |_|
//
// Usage:
// 	%s <command> [arguments]
//
// Use "deeplink help <command>" for more information about a command.
// For more please email hollson@qq.com
// `

var router *gin.Engine

// todo 日志配置文件，
// func InitLog() {
// 	os.MkdirAll("./logs", os.ModePerm)
// 	if out, err := os.Create(fmt.Sprintf("./logs/%s.log", time.Now().Format("20060102150405"))); err != nil {
// 		log.Panic("日志文件创建失败 :", err)
// 	} else {
// 		logrus.SetOutput(out)
// 	}
//
//
// 	fmt.Printf(Welcome, config.App.Version, config.App.Env, config.App.Name)
// }


// 启动程序
func Run() {

	router = gin.Default()
	Route()
	if err := router.Run(config.App.Port); err != nil {
		panic(err)
	}
}
