//-------------------------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.com>
// @ Date: 2019-12-05
// @ Version: 1.0.0
//
// Here's the code description...
//-------------------------------------------------------------------------------------

package config

import (
	"fmt"
	"gopkg.in/gcfg.v1"
)

var App *AppConfig
var Mysql *MysqlConfig
var Redis *RedisConfig

var DbSource string //数据库连接字符串

type AppConfig struct {
	Name string
	Port string
}
type MysqlConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Schema   string
	Charset  string
}

type RedisConfig struct {
	Host string
}

type iniConfig struct {
	App   AppConfig
	Mysql MysqlConfig
	Redis RedisConfig
}

func init() {
	cfg := new(iniConfig)
	err := gcfg.ReadFileInto(cfg, "./conf/config.ini")
	if err != nil {
		panic(err)
	}
	App = &cfg.App
	Mysql = &cfg.Mysql
	Redis = &cfg.Redis

	//"root:123456@(localhost:3306)/testdb?charset=utf8"
	DbSource = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s", Mysql.User, Mysql.Password, Mysql.Host, Mysql.Port, Mysql.Schema, Mysql.Charset)
}
