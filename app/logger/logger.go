package logger

import (
    "io/ioutil"

    "github.com/hollson/gooz/app/config"
    "github.com/hollson/gooz/app/logger/hook"
    "github.com/sirupsen/logrus"
)

// 将7种异常简化为error，info和trance三种
func GetLogLevel(lvl string) int {
    switch lvl {
    case "panic", "fatal", "error":
        return 2
    case "warn", "info":
        return 4
    case "debug", "trance":
        return 6
    default:
        return 4
    }
}

// 从配置文件加载Log配置，并初始化Log Hook
func Init() {
    // 打印行号
    logrus.SetReportCaller(true)

    // 日志级别
    level := logrus.Level(GetLogLevel(config.Log.Level))
    logrus.SetLevel(level)

    // 1. 文本日志输出
    if config.Log.File != nil {
        txt := &hook.TextHookConfig{
            LogPath:      config.Log.File.Path,
            Rotation:     config.Log.File.Rotation,
            Retain:       config.Log.File.Retain,
            LimitSize:    config.Log.File.LimitSize,
            MonitorCycle: config.Log.File.Monitor,
        }

        hk, err := txt.NewHook()
        if err != nil {
            logrus.Errorln("new file hook error:", err)
            return
        }
        logrus.AddHook(hk)
    }

    // 2. Elastic输出
    if config.Log.Elastic != nil {
        es := &hook.EsHook{
            Urls:      config.Log.Elastic.Urls,
            Account:   config.Log.Elastic.Account,
            Password:  config.Log.Elastic.Password,
            IndexName: config.Log.Elastic.Index,
            Level:     level,
        }
        hk, err := es.NewHook()
        if err != nil {
            logrus.Errorln(err)
            return
        }
        if info, status, err := es.PingElastic(); err != nil {
            logrus.Errorf("ping elastic error: info=%v,status=%d,err=%v", info, status, err)
        } else {
            logrus.AddHook(hk)
        }
    }

    // 3. 控制台输出
    if !config.Log.ConsoleOut {
        logrus.SetOutput(ioutil.Discard)
    }
}
