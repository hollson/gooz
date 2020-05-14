//-------------------------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.com>
// @ Date: 2019-12-05
// @ Version: 1.0.0
//
// 按模块定义配置对象
//-------------------------------------------------------------------------------------

package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/gcfg.v1"
)

var App *app           //App配置
var Mysql *mysql       //Mysql数据库
var Redis *redis       //Redis配置
var Postgres *postgres //Postgres数据库
//var Log *Log
//var Zk *Zookeeper
//var etcd *Etcd
//var ConnStr string //拼接的连接字符串

// 运行环境
type Env string
const (
	Env_DEV   Env = "dev"   //开发环境
	Env_TEST  Env = "test"  //测试环境
	Env_STAGE Env = "stage" //验收环境
	Env_PROD  Env = "prod"  //成产环境
)

type app struct {
	Name    string
	Port    string
	Env     Env    //运行环境
	Version string //版本号
}

type mysql struct {
	Enable   bool //是否启用mysql数据库
	Host     string
	Port     string
	User     string
	Password string
	Schema   string
	Charset  string
	ConnStr  string `ini:"-"` //拼接的连接字符串
}

type postgres struct {
	Enable bool   //是否启用mysql数据库
	Source string //链接字符串
}

type redis struct {
	Host string
}

type config struct {
	App      app
	Mysql    mysql
	Postgres postgres
	Redis    redis
}

func init() {
	cfg := config{}
	err := gcfg.ReadFileInto(&cfg, "./conf/app.ini")
	if err != nil {
		logrus.Fatal(err)
	}
	//"root:123456@(localhost:3306)/testdb?charset=utf8"
	cfg.Mysql.ConnStr = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s", cfg.Mysql.User, cfg.Mysql.Password,
		cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.Schema, cfg.Mysql.Charset)

	App = &cfg.App
	Mysql = &cfg.Mysql
	Postgres = &cfg.Postgres
	Redis = &cfg.Redis
}
