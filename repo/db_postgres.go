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
	"github.com/go-xorm/xorm"
	"github.com/hollson/deeplink/etc"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	//"github.com/xormplus/xorm"
	//lg "github.com/xormplus/xorm/log"

	"xorm.io/core"
)

var PG *xorm.Engine
//var Eg *xorm.EngineGroup  //集群模式


func init() {
	if !etc.Postgres.Enable {
		return
	}

	PG, err := xorm.NewEngine("postgres", etc.Postgres.Source)
	if err != nil {
		logrus.Errorln("Postgres Engine错误:", err.Error())
	}

	if err := PG.Ping(); err != nil {
		logrus.Errorln(" ❌ Postgres数据库连接失败:", err)
	} else {
		logrus.Infoln(" ✅ Postgres数据库连接成功 !!!")
	}

	//cluster1 := etc.Postgres.Source
	// 集群模式 Todo:从配置文件读取
	//conns := []string{
	//	cluster1,
	//}
	// 集群访问方式：
	//Eg, err = xorm.NewEngineGroup("postgres", conns, xorm.RandomPolicy()) //随机
	////此时设置的test1数据库和test2数据库的随机访问权重为2和3
	//Eg, err = xorm.NewEngineGroup("postgres", conns, xorm.WeightRandomPolicy([]int{2, 3})) //权重随机
	//Eg, err = xorm.NewEngineGroup("postgres", conns, xorm.RoundRobinPolicy())              //轮询
	//Eg, err = xorm.NewEngineGroup("postgres", conns, xorm.WeightRoundRobinPolicy([]int{2, 3})) //权重轮询
	//Eg, err = xorm.NewEngineGroup("postgres", conns, xorm.LeastConnPolicy()) //最小链接

	if etc.App.Env == etc.Env_PROD {
		PG.ShowSQL(false)
		PG.SetLogLevel(core.LOG_ERR)
		//PG.SetMaxIdleConns(30) 	//最大空闲数
		//PG.SetMaxOpenConns(500)	//最大连接数
	} else {
		PG.ShowSQL(true)
		PG.SetLogLevel(core.LOG_INFO)
	}
}
