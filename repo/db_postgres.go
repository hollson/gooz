//-------------------------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.com>
// @ Date: 2020-05-06
// @ Version: 1.0.0
//
// Here's the code description...
//-------------------------------------------------------------------------------------

package repo

import (
	"github.com/hollson/deeplink/etc"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/xormplus/xorm"
	lg "github.com/xormplus/xorm/log"
)

var Eg *xorm.EngineGroup

func init() {
	var err error
	cluster1 := etc.Postgres.Source

	// 集群模式 Todo:从配置文件读取
	conns := []string{
		cluster1,
	}

	// 集群访问方式：
	//Eg, err = xorm.NewEngineGroup("postgres", conns, xorm.RandomPolicy()) //随机
	////此时设置的test1数据库和test2数据库的随机访问权重为2和3
	//Eg, err = xorm.NewEngineGroup("postgres", conns, xorm.WeightRandomPolicy([]int{2, 3})) //权重随机
	//Eg, err = xorm.NewEngineGroup("postgres", conns, xorm.RoundRobinPolicy())              //轮询
	//Eg, err = xorm.NewEngineGroup("postgres", conns, xorm.WeightRoundRobinPolicy([]int{2, 3})) //权重轮询
	Eg, err = xorm.NewEngineGroup("postgres", conns, xorm.LeastConnPolicy()) //最小链接
	if err != nil {
		logrus.Errorln("Postgres Engine错误:", err.Error())
	}

	if err := Eg.Ping(); err != nil {
		logrus.Errorln(" ❌ Postgres数据库连接失败:", err)
	} else {
		logrus.Infoln(" ✅ Postgres数据库连接成功 !!!")
	}

	if etc.App.Env == etc.Env_PROD {
		DB.ShowSQL(false)
		DB.SetLogLevel(lg.LOG_ERR)
		//DB.SetMaxIdleConns(30) 	//最大空闲数
		//DB.SetMaxOpenConns(500)	//最大连接数
	} else {
		DB.ShowSQL(true)
		DB.SetLogLevel(lg.LOG_INFO)
	}
}
