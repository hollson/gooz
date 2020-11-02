// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package hook

import (
    "github.com/sirupsen/logrus"
)

type Hook interface {
    NewHook() (logrus.Hook, error)
}

//
// import (
//     "io/ioutil"
//
//     "github.com/sirupsen/logrus"
// )
//
// const (
//     LogLevel    = "error" // 日志级别 (error,info或trance)
//     ConsoleOut  = false   // 控制台输出
//     ElasticOut  = false   // Es日志输出
//     TextFileOut = false   // 文件日志输出
// )
//
// // 将7种异常简化为error，info和trance三种
// func GetLogLevel(lvl string) int {
//     switch lvl {
//     case "panic", "fatal", "error":
//         return 2
//     case "warn", "info":
//         return 4
//     case "debug", "trance":
//         return 6
//     default:
//         return 4
//     }
// }
//

//
// // 从配置文件加载Log配置，并初始化Log Hook
// func Init() {
//     // 打印行号
//     logrus.SetReportCaller(true)
//
//     // 日志级别
//     level := logrus.Level(GetLogLevel(LogLevel))
//     logrus.SetLevel(level)
//
//     // 1. 文本日志输出
//     if TextFileOut {
//         texthook := &TextHookConfig{
//             LogPath:      "/Users/sybs/logs/product",
//             Rotation:     1,
//             Retain:       1,
//             LimitSize:    200,
//             MonitorCycle: 10,
//         }
//
//         hook, err := texthook.NewHook()
//         if err != nil {
//             logrus.Errorln("new file hook error:", err)
//             return
//         }
//         logrus.AddHook(hook)
//     }
//
//     // 2. Elastic输出
//     if ElasticOut {
//         hk := &EsHook{
//             Urls:      []string{"http://10.0.0.13:9200"},
//             Account:  "elastic",
//             Password: "123456",
//             IndexName: "reworld-server",
//             Level:     level,
//         }
//         hook, err := hk.NewHook()
//         if err != nil {
//             logrus.Errorln("new elastic hook error:", err)
//             return
//         }
//         if _,_,err:=hk.PingElastic();err==nil{
//             logrus.AddHook(hook)
//         }
//     }
//
//     // 3. 控制台输出
//     if !ConsoleOut {
//         logrus.SetOutput(ioutil.Discard)
//     }
// }
