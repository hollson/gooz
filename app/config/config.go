//-------------------------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.cn>
// @ Date: 2019-12-05
// @ Version: 1.0.0
//
// 解析app.toml配置文件
// 使用示例：https://blog.csdn.net/Gusand/article/details/106094535
//-------------------------------------------------------------------------------------

package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/hollson/deeplink/util"
	"github.com/sirupsen/logrus"
	"os"
	"time"
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

//参考Viki：https://en.wikipedia.org/wiki/Deployment_environment
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

func main() {
	file, _ := os.Create("d:/test.log") //创建文件
	defer file.Close()

	num, _ := file.Write([]byte("hello"))
	fmt.Printf("写入 %d 个字节n", num)
}

func init() {
	var cfg config
	//todo 命令行可创建模板配置文件

	//按照./app.config和./conf/app.toml目录优先级加载配置文件，都不存在时使用模板创建配置文件。
	if pth := "./app.toml"; util.Exists(pth) {
		if _, err := toml.DecodeFile(pth, &cfg); err != nil {
			logrus.Panic(err)
		}
	} else if pth:= "./conf/app.toml"; util.Exists(pth) {
		if _, err := toml.DecodeFile(pth, &cfg); err != nil {
			logrus.Panic(err)
		}
	} else {
		util.CreateFile("./conf")
		f, err := os.Create(pth)
		defer f.Close()
		if err != nil {
			logrus.Panic(err)
		}
		f.WriteString(fmt.Sprintf(tpl,time.Now().Format("2006-01-02 15:04:05")))
		logrus.Infof(" 👷 初始化配置文件创建成功！！！")

		if _, err := toml.Decode(tpl, &cfg); err != nil {
			panic(err)
		}
	}

	//Mysql链接字符串："user:pwd@(host:port)/dbname?charset=utf8"
	cfg.Mysql.Source = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s", cfg.Mysql.User,
		cfg.Mysql.Password, cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.Schema, cfg.Mysql.Charset)

	//Postgres链接字符串："postgres://user:pwd@host:port/dbname?sslmode=disable;"
	for k, v := range cfg.Postgres {
		cfg.Postgres[k].Source = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s;",
			v.User, v.Password, v.Host, v.Port, v.Schema, v.Sslmode)
	}

	App = &cfg.App
	Log = &cfg.Log
	Mysql = &cfg.Mysql
	Postgres = &cfg.Postgres
	Redis = cfg.Redis
}
