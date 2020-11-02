// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package hook

import (
    "time"

    "github.com/hollson/goox/memory"
    "github.com/sirupsen/logrus"
)

// 人性化(配置文件)Elastic Hook对象，用于生产
type TextHookConfig struct {
    LogPath      string // 日志根目录,默认：./logs
    Rotation     uint32 // 文件切割周期(天),默认：1天
    Retain       uint32 // 日志保留时间(天)，默认：7天
    LimitSize    uint32 // 单个文件大小上限(M)，默认：100M
    MonitorCycle uint32 // 文件监控周期(秒)，默认：5s
}

// 是TextHook简化版，适用于配置文件形式的Hook实例化
func (h *TextHookConfig) NewHook() (logrus.Hook, error) {
    if h.LogPath == "" {
        h.LogPath = "./logs"
    }
    if h.Rotation < 1 {
        h.Rotation = 1
    }
    if h.Retain < 1 {
        h.Retain = 7
    }
    if h.LimitSize < 1 {
        h.LimitSize = 100
    }
    if h.MonitorCycle < 1 {
        h.MonitorCycle = 5
    }
    hook := &TextHook{
        LogPath:      h.LogPath,
        CutCycle:     time.Duration(h.Rotation) * time.Hour * 24,
        RetainTime:   time.Duration(h.Retain) * time.Hour * 24,
        FileLimit:    memory.Size(h.LimitSize) * memory.MB,
        MonitorCycle: time.Duration(h.MonitorCycle) * time.Second,
    }
    return hook.NewHook()
}
