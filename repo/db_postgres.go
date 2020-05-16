// -------------------------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.cn>
// @ Date: 2020-05-06
// @ Version: 1.0.0
//
// Here's the code description...
// -------------------------------------------------------------------------------------

package repo

import (
	"fmt"
	"strings"

	"github.com/hollson/deeplink/app/config"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/xormplus/xorm"
	lg "github.com/xormplus/xorm/log"
)

// var PG *xorm.Engine
var PG *xorm.EngineGroup // 集群模式

func init() {
	// 添加集群连接字符串
	var conns, tips []string
	for _, val := range *config.Postgres {
		if val.Enable {
			pg, err := xorm.NewEngine("postgres", val.Source)
			if err != nil {
				logrus.Errorf("Postgres Engine错误:%v", err.Error())
				continue
			}
			if err := pg.Ping(); err != nil {
				logrus.Errorf(" ❌  Postgres【%s:%d/%s】连接失败,%v", val.Host, val.Port, val.Schema, err)
				continue
			}
			tips = append(tips, fmt.Sprintf("【%s:%d/%s】", val.Host, val.Port, val.Schema))
			conns = append(conns, val.Source)
			logrus.Infof(" ✅  Postgres【%s:%d/%s】连接成功", val.Host, val.Port, val.Schema)
		}
	}

	if len(conns) < 1 {
		return
	}
	logrus.Infof(" 🍇 Postgres集群状态：%v️", strings.Join(tips, ""))

	// 集群一种集群访问方式：
	var err error
	PG, err = xorm.NewEngineGroup("postgres", conns, xorm.RandomPolicy()) // 随机
	// 此时设置的test1数据库和test2数据库的随机访问权重为2和3
	// PG, err = xorm.NewEngineGroup("postgres", conns, xorm.WeightRandomPolicy([]int{2, 3})) //权重随机
	// PG, err = xorm.NewEngineGroup("postgres", conns, xorm.RoundRobinPolicy())              //轮询
	// PG, err = xorm.NewEngineGroup("postgres", conns, xorm.WeightRoundRobinPolicy([]int{2, 3})) //权重轮询
	// PG, err = xorm.NewEngineGroup("postgres", conns, xorm.LeastConnPolicy()) //最小链接
	if err != nil {
		logrus.Errorln(err.Error())
		return
	}

	PG.ShowSQL(true)
	PG.SetLogLevel(lg.LOG_ERR)
	if config.App.Env == config.Env_PROD {
		PG.ShowSQL(false)
		PG.SetLogLevel(lg.LOG_ERR)
		// PG.SetMaxIdleConns(30) 	//最大空闲数
		// PG.SetMaxOpenConns(500)	//最大连接数
	}
}
