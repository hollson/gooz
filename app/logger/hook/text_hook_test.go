// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package hook

import (
    "testing"
    "time"

    "github.com/hollson/goox/memory"
    "github.com/sirupsen/logrus"
)

func TestTextHook(t *testing.T) {
    logrus.SetLevel(logrus.WarnLevel)
    logrus.SetReportCaller(true) // 打印行号
    texthook := &TextHook{
        LogPath:      "/Users/sybs/logs/user",
        CutCycle:     time.Hour * 24,
        RetainTime:   time.Hour * 24 * 7,
        FileLimit:    memory.KB * 10,
        MonitorCycle: time.Second * 5,
    }
    hook, err := texthook.NewHook()
    if err != nil {
        logrus.Errorln("hook.NewHook error:", err)
        return
    }
    logrus.AddHook(hook)

    for {
        logrus.Errorln("出现了异常，赶紧去处理吧")
        logrus.Infoln("恭喜你，这段代码是正常的哦 恭喜你，这段代码是正常的哦")
        time.Sleep(time.Second * 1)
    }
}

func TestTextConfigtHook(t *testing.T) {
    logrus.SetLevel(logrus.WarnLevel)
    logrus.SetReportCaller(true) // 打印行号
    texthook := &TextHookConfig{
        LogPath:      "/Users/sybs/logs/product",
        Rotation:     1,
        Retain:       1,
        LimitSize:    200,
        MonitorCycle: 10,
    }

    hook, err := texthook.NewHook()
    if err != nil {
        logrus.Errorln("hook.NewHook error:", err)
        return
    }
    logrus.AddHook(hook)

    for {
        logrus.Errorln("配置文件，出现了异常，赶紧去处理吧")
        logrus.Infoln("配置文件，恭喜你，成功了")
        time.Sleep(time.Second * 1)
    }
}

func TestEsHook(t *testing.T) {
    logrus.SetLevel(logrus.InfoLevel)
    logrus.SetReportCaller(true) // 打印行号
    hk := &EsHook{
        Urls:      []string{"http://10.0.0.13:9200"},
        Account:   "elastic",
        Password:  "123456",
        IndexName: "reworld-server",
        Level:     logrus.InfoLevel,
    }

    hook, err := hk.NewHook()
    if err != nil {
        logrus.Errorln("hook.NewHook error:", err)
        return
    }
    logrus.AddHook(hook)

    for {
        logrus.Errorf("%d,出现了异常", time.Now().Second())
        logrus.Infof("%d,恭喜你，成功了", time.Now().Second())
        time.Sleep(time.Second * 3)
    }
}
