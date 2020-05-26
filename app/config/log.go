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
	"errors"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/hollson/gox/file"
	"github.com/hollson/gox/memory"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// todo 配置文件
const (
	LogRoot                = "./logs"
	TIDYPATH          bool = false // 精简文件路径
	MAXLOGSIZE             = 200   // 文件最大尺寸
	INFOLEVEL              = 4     // 1-6  ,默认4
	DISABLEHTMLESCAPE      = true

	ESHOOK_ENABLE     = false
	INFLUXHOOK_ENABLE = false
)

var (
	ROTATIONTIME = time.Hour * 24     // 日志周期(默认每86400秒/一天旋转一次)
	WITHMAXAGE   = time.Hour * 24 * 7 // 默认值：-1
)

// 设置日志规则
// todo 配置hook时，无须log文件
func InitLog() {
	logrus.SetLevel(INFOLEVEL) // 日志级别
	if hk, err := MultiSplitHook(""); err != nil {
		logrus.Errorf(" ❌ Local log hook error,%v", err)
	} else {
		logrus.Infof(" 🐸 Local hook OK")
		logrus.AddHook(hk)
	}

	if hk, err := ESHook(""); err != nil {
		logrus.Errorf(" ❌ Elastic search log hook error,%v", err)
	} else {
		logrus.Infof(" 🐸 Es hook OK")
		logrus.AddHook(hk)
	}

	if hk, err := InfluxDBHook(""); err != nil {
		logrus.Errorf(" ❌  InfluxDb log hook error,%v", err)
	} else {
		logrus.Infof(" 🐸 InfluxDB hook OK")
		logrus.AddHook(hk)
	}

	logrus.SetReportCaller(true) // 打印行号
	// logrus.SetOutput(ioutil.Discard) // 关闭日志输出
}

// 本地日志多文件切分Hook
func MultiSplitHook(prefix string) (hk logrus.Hook, err error) {
	if len(prefix) > 0 {
		prefix = prefix + "-"
	}
	var accessWriter, errorWriter *rotatelogs.RotateLogs

	// logs/deeplink-20200203-access.log
	accessWriter, err = rotatelogs.New(
		fmt.Sprintf("%s/%s%s-access.log", LogRoot, prefix, "%Y%m%d"),
		rotatelogs.ForceNewFile(),
		rotatelogs.WithRotationTime(ROTATIONTIME),
		rotatelogs.WithMaxAge(WITHMAXAGE),
	)

	if err != nil {
		return nil, err
	}

	// logs/deeplink-20200203-error.log
	errorWriter, err = rotatelogs.New(
		fmt.Sprintf("%s/%s%s-error.log", LogRoot, prefix, "%Y%m%d"),
		rotatelogs.ForceNewFile(),
		rotatelogs.WithRotationTime(ROTATIONTIME),
		rotatelogs.WithMaxAge(WITHMAXAGE),
	)
	if err != nil {
		return nil, err
	}

	// 对Access文件进行文件大小切分
	if MAXLOGSIZE > memory.KB*128 {
		go func() {
			runtime.Gosched()
			tk := time.Tick(time.Second * 10)
			for {
				<-tk
				cf := accessWriter.CurrentFileName()
				size := file.FileSize(cf)
				if size > int64(memory.MB)*MAXLOGSIZE {
					logrus.Infoln("Log file size reaches maximum limit :", size)
					if err := accessWriter.Rotate(); err != nil {
						logrus.Errorln(err)
						continue
					}
				}
			}
		}()
	}

	// 可设置按不同level创建不同的文件名
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: accessWriter,
		logrus.InfoLevel:  accessWriter,

		logrus.WarnLevel:  errorWriter,
		logrus.ErrorLevel: errorWriter,
		logrus.FatalLevel: errorWriter,
		logrus.PanicLevel: errorWriter,
	}, jsonFormat())

	return lfsHook, nil
}

// InfluxDb Hook
func InfluxDBHook(prefix string) (logrus.Hook, error) {
	if !INFLUXHOOK_ENABLE {
		return nil, errors.New("influxDb hook can not be enabled")
	}
	return nil, nil
}

// Es hook
func ESHook(prefix string) (logrus.Hook, error) {
	if !ESHOOK_ENABLE {
		return nil, errors.New("es hook can not be enabled")
	}
	return nil, nil
}

func jsonFormat() *logrus.JSONFormatter {
	return &logrus.JSONFormatter{
		PrettyPrint:       false,
		DisableHTMLEscape: DISABLEHTMLESCAPE,
		TimestampFormat:   "2006-01-02 15:04:05",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			if TIDYPATH {
				if frame == nil || frame.Func == nil {
					return "runtime.Frame.Function.Init", "runtime.Frame.File Init"
				}
				fn := strings.LastIndex(frame.Function, "/")
				_, ln := frame.Func.FileLine(frame.PC)
				fi := strings.LastIndex(frame.File, "/")
				return frame.Function[fn+1:], fmt.Sprintf("%s:%d", frame.File[fi+1:], ln)
			}
			return
		},
	}
}
