// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import (
    "fmt"
)

type app struct {
    Name    string
    Port    string
    Env     Env    `toml:"environment"` // 运行环境
    Version string // 版本号
}

type logger struct {
    Level      string `toml:"level"`   // 日志级别(error => info => trance)
    ConsoleOut bool   `toml:"console"` // 是否控制台输出

    // 文件日志输出
    File *struct {
        Path      string `toml:"path"`       // 日志输出路径
        Rotation  uint32 `toml:"rotation"`   // 文件切割周期(天),默认：1天
        Monitor   uint32 `toml:"monitor"`    // 文件监控周期(秒)，默认：5s
        Retain    uint32 `toml:"retain"`     // 日志保留时间(天)，默认：7天
        LimitSize uint32 `toml:"limit_size"` // 单个文件大小上限(M)，默认：100M
    }

    Elastic *struct {
        Urls     []string `toml:"urls"`     // es集群
        Account  string   `toml:"account"`  // 账号
        Password string   `toml:"password"` // 密码
        Index    string   `toml:"index"`    // 索引名称
    }
}

type mysql struct {
    Enable   bool // 是否启用mysql数据库
    Host     string
    Port     int
    User     string
    Password string
    Schema   string
    Charset  string
}

// 链接字符串：
// "user:pwd@(host:port)/dbname?charset=utf8"
func (p *mysql) String() string {
    if len(p.Charset) == 0 {
        p.Charset = "utf8"
    }
    return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s",
        p.User, p.Password, p.Host, p.Port, p.Schema, p.Charset)
}

type postgres struct {
    Enable   bool
    Host     string
    Port     int
    User     string
    Password string
    Schema   string
    Sslmode  string
}

// 链接字符串：
// "postgres://user:pwd@host:port/dbname?sslmode=disable;"
func (p *postgres) String() string {
    if len(p.Sslmode) == 0 {
        p.Sslmode = "disable"
    }
    return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s;",
        p.User, p.Password, p.Host, p.Port, p.Schema, p.Sslmode)
}

type redis struct {
    Host string
    Port int
    Auth string
}

type sentinel struct {
    Psssworld  string
    MasterName string
    Member     []redis
}

type cluster struct {
    Member []redis
}
