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

	"github.com/hollson/deeplink/etc"
	"github.com/sirupsen/logrus"
	"github.com/xormplus/xorm"
	lg "github.com/xormplus/xorm/log"
)

//定义orm引擎
var DB *xorm.Engine

//创建orm引擎
func init() {
	InitMysql()
}

func InitMysql() {
	if !etc.Mysql.Enable {
		return
	}

	var err error
	DB, err = xorm.NewEngine("mysql", etc.Mysql.ConnStr)
	if err != nil {
		logrus.Errorln("Mysql Engine错误:", err.Error())
	} else {
		if err := DB.Ping(); err != nil {
			logrus.Errorln(" ❌ Mysql数据库连接失败:", err)
		} else {
			logrus.Infoln(" ✅ Mysql数据库连接成功 !!!")
		}
	}

	if etc.App.Env == etc.Env_PROD {
		DB.ShowSQL(false)
		DB.SetLogLevel(lg.LOG_ERR)
		//DB.SetMaxIdleConns(30) //最大空闲数
		//DB.SetMaxOpenConns(500)  //最大连接数
	} else {
		DB.ShowSQL(true)
		DB.SetLogLevel(lg.LOG_INFO)
	}
}
