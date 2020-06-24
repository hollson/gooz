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
	Env     Env `toml:"environment"` // 运行环境
	Version string                   // 版本号
}

type logger struct {
	Path  string
	Level string
	Hook  string
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
	if len(p.Charset)==0{
		p.Charset="utf8"
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
	if len(p.Sslmode)==0{
		p.Sslmode="disable"
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
	Member []redis
}

type cluster struct {
	Member []redis
}
