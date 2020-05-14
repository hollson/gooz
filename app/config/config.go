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
	"github.com/BurntSushi/toml"
)

var App *app               //App配置
var Log *log               //日志配置
var Mysql *mysql           //Mysql数据库
var Redis map[string]redis //Redis配置
var Postgres *[]postgres   //Postgres数据库
//var Log *Log
//var Zk *Zookeeper
//var etcd *Etcd

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
	Env     Env `toml:"environment"` //运行环境
	Version string                   //版本号
}

type log struct {
	Path  string
	Level string
	Hook  string
}

type mysql struct {
	Enable   bool //是否启用mysql数据库
	Host     string
	Port     int
	User     string
	Password string
	Schema   string
	Charset  string
	Source   string `toml:"-"` //拼接的连接字符串
}

type postgres struct {
	Enable   bool
	Host     string
	Port     int
	User     string
	Password string
	Schema   string
	Sslmode  string
	Source   string `toml:"-"` //拼接的连接字符串
}

type redis struct {
	Host string
	Port int
}

// 组合
type config struct {
	App      app
	Log      log
	Mysql    mysql
	Postgres []postgres
	Redis    map[string]redis
}

func init() {
	var cfg config
	if _, err := toml.DecodeFile("./conf/app.toml", &cfg); err != nil {
		panic(err)
	}

	//Mysql链接字符串："user:pwd@(host:port)/dbname?charset=utf8"
	cfg.Mysql.Source = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s", cfg.Mysql.User,
		cfg.Mysql.Password, cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.Schema, cfg.Mysql.Charset)

	//Postgres链接字符串："postgres://user:pwd@host:port/dbname?sslmode=disable;"
	for k, v := range cfg.Postgres {
		cfg.Postgres[k].Source =fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s;",
			v.User,v.Password,v.Host,v.Port,v.Schema,v.Sslmode)
	}

	App = &cfg.App
	Log = &cfg.Log
	Mysql = &cfg.Mysql
	Postgres = &cfg.Postgres
	Redis = cfg.Redis
}
