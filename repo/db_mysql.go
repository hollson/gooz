//-------------------------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.com>
// @ Date: 2020-05-06
// @ Version: 1.0.0
//
// Here's the code description...
// 参考教程： https://www.kancloud.cn/xormplus/xorm/167078
//-------------------------------------------------------------------------------------

package repo

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/hollson/deeplink/app/config"
	"github.com/sirupsen/logrus"
	"github.com/xormplus/xorm"
	lg "github.com/xormplus/xorm/log"
)

//定义orm引擎
var My *xorm.Engine

//创建orm引擎
func init() {
	InitMysql()
}

func InitMysql() {
	if !config.Mysql.Enable {
		return
	}

	var err error
	My, err = xorm.NewEngine("mysql", config.Mysql.Source)
	if err != nil {
		logrus.Errorln("Mysql Engine错误:", err.Error())
	} else {
		if err := My.Ping(); err != nil {
			logrus.Errorln(" ❌  Mysql数据库连接失败:", err)
		} else {
			logrus.Infoln(" ✅  Mysql数据库连接成功")
		}
	}

	My.ShowSQL(true)
	My.SetLogLevel(lg.LOG_INFO)
	if config.App.Env == config.Env_PROD {
		My.ShowSQL(false)
		My.SetLogLevel(lg.LOG_ERR)
		//My.SetMaxIdleConns(30) //最大空闲数
		//My.SetMaxOpenConns(500)  //最大连接数
	}
}
