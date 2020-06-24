package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/hollson/deeplink/util"
)

var App *app               // App配置
var Log *logger            // 日志配置
var Mysql *mysql           // Mysql数据库
// var Redis map[string]redis // Redis配置
var Postgres *[]postgres   // Postgres数据库
// var Log *Log
// var Zk *Zookeeper
// var etcd *Etcd

// 全局配置对象
type config struct {
	App      app `toml:"app"`
	Log      logger
	Mysql    mysql
	Postgres []postgres
	// Redis    map[string]redis
}

func load() {
	var cfg config
	// todo 命令行可创建模板配置文件
	// todo 如果有服务注册，则分两步，读etcd或zk配置，再加载相关配置

	// 按照./app.config和./conf/app.toml目录优先级加载配置文件，都不存在时使用模板创建配置文件。
	if pth := "./config.toml"; util.Exists(pth) {
		if _, err := toml.DecodeFile(pth, &cfg); err != nil {
			log.Panic(err)
		}
	} else if pth := "./conf/config.toml"; util.Exists(pth) {
		if _, err := toml.DecodeFile(pth, &cfg); err != nil {
			log.Panic(err)
		}
	} else {
		util.CreateFile("./conf")
		f, err := os.Create(pth)
		defer f.Close()
		if err != nil {
			log.Panic(err)
		}
		f.WriteString(fmt.Sprintf(TPL, time.Now().Format("2006-01-02 15:04:05")))
		log.Println(" 👷 初始化配置文件创建成功！！！")

		if _, err := toml.Decode(TPL, &cfg); err != nil {
			panic(err)
		}
	}

	App = &cfg.App
	Log = &cfg.Log
	Mysql = &cfg.Mysql
	Postgres = &cfg.Postgres
	// Redis = cfg.Redis
}

func Init() {
	load()
}
