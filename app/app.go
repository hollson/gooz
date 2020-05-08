//-------------------------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.com>
// @ Date: 2019-12-01
// @ Version: 1.0.0
//-------------------------------------------------------------------------------------

package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hollson/deeplink/etc"
	_ "github.com/hollson/deeplink/repo"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

//http://www.network-science.de/ascii/
const FAVORITE = `
Deeplink is a internal applications <XX科技.保留版权>
       _              _ _      _
    __| |___ ___ _ __| (_)_ _ | |__
   / _  / -_) -_) '_ \ | | ' \| / /
   \____\___\___| .__/_|_|_||_|_\_\   %s(%s)
                |_|

Usage:
	%s <command> [arguments] 

Use "deeplink help <command>" for more information about a command.
For more please email hollson@qq.com
`

var router *gin.Engine

// 设置日志规则
func init() {
	fmt.Printf(FAVORITE,etc.App.Version,etc.App.Env,etc.App.Name)

	os.MkdirAll("./logs",os.ModePerm)
	if log, err := os.Create(fmt.Sprintf("./logs/%s.log", time.Now().Format("20060102150405"))); err != nil {
		logrus.Errorln("Create file err :", err)
	} else {
		logrus.SetOutput(log)
	}
}

// 启动程序
func Run() {
	router = gin.Default()
	Route()
	if err := router.Run(etc.App.Port); err != nil {
		panic(err)
	}
}
