// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

/**************************************
  文本文件格式的日志输出
**************************************/
package hook

import (
    "fmt"
    "runtime"
    "strings"
    "time"

    "github.com/hollson/goox/file"
    "github.com/hollson/goox/memory"
    rotate "github.com/lestrrat-go/file-rotatelogs"
    "github.com/rifflock/lfshook"
    "github.com/sirupsen/logrus"
)

const (
    HTMLESCAPE = false     // htmlEscape:是否对html标签进行转义
    TIDYPATH   = true      // tidyPath:是否简化函数与文件的输出路径
    ACC_PREFIX = "access-" // 访问日志前缀
    ERR_PREFIX = "error-"  // 错误日志前缀
)

// 原生Elastic Hook对象
type TextHook struct {
    LogPath      string        // 日志根目录
    CutCycle     time.Duration // 文件切割周期，如按日期新建日志文件
    RetainTime   time.Duration // 日志保留时间
    FileLimit    memory.Size   // 单个文件大小上限
    MonitorCycle time.Duration // 文件监控周期,即定期检查文件是否达到上限
}

func (h *TextHook) NewHook() (logrus.Hook, error) {
    // === 输出并监控访问日志 ===
    accWriter, err := h.newRotate(ACC_PREFIX)
    if err != nil {
        logrus.Errorln("create accessWriter error", err)
    }
    h.monitorLog(accWriter)

    // === 输出并监控错误日志 ===
    errWriter, err := h.newRotate(ERR_PREFIX)
    if err != nil {
        logrus.Errorln("create errorWriter error", err)
    }
    h.monitorLog(errWriter)

    // 日志钩子
    lfsHook := lfshook.NewHook(lfshook.WriterMap{
        logrus.DebugLevel: accWriter,
        logrus.InfoLevel:  accWriter,

        logrus.WarnLevel:  errWriter,
        logrus.ErrorLevel: errWriter,
        logrus.FatalLevel: errWriter,
        logrus.PanicLevel: errWriter,
    }, jsonFormat(HTMLESCAPE, TIDYPATH))
    return lfsHook, nil
}

// 创建日志文件旋转控制器
// path: 日志目录
// prefix：文件名前缀，如：给202001.log文件加前缀，变成access-202001.log
func (h *TextHook) newRotate(prefix string) (*rotate.RotateLogs, error) {
    var accessWriter *rotate.RotateLogs
    accessWriter, err := rotate.New(
        fmt.Sprintf("%s/%s%s.log", h.LogPath, prefix, "%Y%m%d"),
        rotate.WithRotationTime(h.CutCycle),
        rotate.WithMaxAge(h.RetainTime),
    )
    if err != nil {
        return nil, err
    }
    return accessWriter, nil
}

// (异步)监控日志文件大小，到达上限时，强制旋转
func (h *TextHook) monitorLog(writer *rotate.RotateLogs) {
    go func() {
        runtime.Gosched()
        tk := time.Tick(h.MonitorCycle)
        for {
            <-tk
            cf := writer.CurrentFileName()
            size := memory.Size(file.FileSize(cf))
            if size > h.FileLimit {
                logrus.Infoln("log file size reaches maximum limit :", size)
                if err := writer.Rotate(); err != nil {
                    logrus.Errorln(err)
                    continue
                }
            }
        }
    }()
}

// htmlEscape:是否对html标签进行转义
// tidyPath:是否简化函数与文件的输出路径
func jsonFormat(htmlEscape, tidyPath bool) *logrus.JSONFormatter {
    return &logrus.JSONFormatter{
        PrettyPrint:       false,       // 不格式化(美化)
        DisableHTMLEscape: !htmlEscape, // 是否对html标签进行转义
        TimestampFormat:   "2006-01-02 15:04:05",
        CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
            if frame == nil || frame.Func == nil {
                return "runtime.Frame.Function.Init", "runtime.Frame.File Init"
            }
            if tidyPath {
                fn := strings.LastIndex(frame.Function, "/")
                _, ln := frame.Func.FileLine(frame.PC)
                fi := strings.LastIndex(frame.File, "/")
                return frame.Function[fn+1:], fmt.Sprintf("%s:%d", frame.File[fi+1:], ln)
            }
            return frame.Function, frame.File
        },
    }
}
