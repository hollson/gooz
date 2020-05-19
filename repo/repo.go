// -------------------------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.cn>
// @ Date: 2020-05-07
// @ Version: 1.0.0
//
// Here's the code description...
// -------------------------------------------------------------------------------------

package repo

import (
	"github.com/garyburd/redigo/redis"
	"xorm.io/xorm"
)

// Mysql数据引擎
var My *xorm.Engine


// Postgres数据引擎
// var PG *xorm.Engine      // 单机模式
var PG *xorm.EngineGroup // 集群模式


// Redis数据库
var Rds []redis.Conn