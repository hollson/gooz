// -------------------------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.cn>
// @ Date: 2020-05-06
// @ Version: 1.0.0
//
// Here's the code description...
// 参考教程： https://www.kancloud.cn/xormplus/xorm/167078
// -------------------------------------------------------------------------------------

package repo

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/hollson/deeplink/app/config"
	"github.com/sirupsen/logrus"
	"xorm.io/xorm"
	xlg "xorm.io/xorm/log"
)

func InitMysql() {
	if !config.Mysql.Enable {
		return
	}

	var err error
	My, err = xorm.NewEngine("mysql", config.Mysql.Source)
	if err != nil {
		logrus.Errorln(" ❌  Mysql Engine错误:", err.Error())
		return
	}

	if err := My.Ping(); err != nil {
		logrus.Errorln(" ❌  Mysql数据库连接失败:", err)
		return
	}
	logrus.Infoln(" ✅  Mysql数据库连接成功")

	// todo 配置
	My.ShowSQL(true)
	My.SetLogLevel(xlg.LOG_INFO)
	if config.App.Env == config.Env_PROD {
		My.ShowSQL(false)
		My.SetLogLevel(xlg.LOG_ERR)
		// My.SetMaxIdleConns(30) //最大空闲数
		// My.SetMaxOpenConns(500)  //最大连接数
	}
}
