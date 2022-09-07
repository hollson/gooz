package main

import (
    "runtime"
    "time"

    "github.com/hollson/gooz/app/config"
    "github.com/hollson/gooz/app/logger"
    "github.com/hollson/gooz/repo"
    "github.com/hollson/gooz/router"
    "github.com/sirupsen/logrus"
)

// 按顺序加载初始化项
func init() {
    config.Init()
    router.Init()
    repo.Init()
    logger.Init()
}
//https://threedots.tech/post/introducing-clean-architecture/

//go:generate go build -o gooz
func main() {
    go func() {
        for {
            logrus.Errorf("%d,抱歉，出现了异常", time.Now().Second())
            logrus.Infof("%d,可喜可贺，成功了", time.Now().Second())
            time.Sleep(time.Second * 3)
        }
    }()

    runtime.Gosched()
    router.Run()
}
