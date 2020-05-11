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
	"github.com/hollson/deeplink/app/config"
	_ "github.com/hollson/deeplink/repo"
	logredis "github.com/rogierlommers/logrus-redis-hook"
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
	fmt.Printf(FAVORITE, config.App.Version, config.App.Env, config.App.Name)

	//todo 日志配置文件
	os.MkdirAll("./logs", os.ModePerm)
	if log, err := os.Create(fmt.Sprintf("./logs/%s.log", time.Now().Format("20060102150405"))); err != nil {
		logrus.Errorln("Create file err :", err)
	} else {
		logrus.SetOutput(log)
	}

	logrus.SetLevel(logrus.InfoLevel) //日志级别
	logrus.SetReportCaller(true)      //打印行号

	//Json输出
	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint:     false,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 设置Hook
	hookConfig := logredis.HookConfig{
		Host:     "localhost",
		Key:      "test",
		Format:   "v1",
		App:      "my_app_name",
		Port:     6379,
		Hostname: "my_app_hostname",
		DB:       0, // optional
		TTL:      3600,
	}
	hook, err := logredis.NewHook(hookConfig)
	if err == nil {
		logrus.AddHook(hook)
	} else {
		logrus.Errorf("logredis error: %q", err)
	}

	// 公共信息，可组合
	logrus.WithFields(logrus.Fields{"module": "user"}).
		WithFields(logrus.Fields{"node": "master"}).
		Infoln("日志内容")

	//关闭日志输出
	//logrus.SetOutput(ioutil.Discard)
}

// 启动程序
func Run() {
	router = gin.Default()
	Route()
	if err := router.Run(config.App.Port); err != nil {
		panic(err)
	}
}
