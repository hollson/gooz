// -----------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.cn>
// @ Date: 2020-05-18
// @ Version: 1.0.0
//
// Here's the code description...
// -----------------------------------------------------------------------

package config

import (
	"fmt"
	"log"
	"os"
	"time"

	logredis "github.com/rogierlommers/logrus-redis-hook"
	"github.com/sirupsen/logrus"
)

// 设置日志规则
// todo 配置hook时，无须log文件
func InitLog() {
	os.MkdirAll("./logs", os.ModePerm)
	if out, err := os.Create(fmt.Sprintf("./logs/%s.log", time.Now().Format("20060102150405"))); err != nil {
		log.Panic("日志文件创建失败 :", err)
	} else {
		logrus.SetOutput(out)
	}

	logrus.SetLevel(logrus.InfoLevel) // 日志级别
	logrus.SetReportCaller(true)      // 打印行号

	// Json输出
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
		log.Print("logredis error: %q", err)
	}

	// 公共信息，可组合
	logrus.WithFields(logrus.Fields{"module": "user"}).
		WithFields(logrus.Fields{"node": "master"}).
		Infoln("日志内容")

	// 关闭日志输出
	// logrus.SetOutput(ioutil.Discard)
}
