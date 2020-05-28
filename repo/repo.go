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

func Init() {
	initMysql()    // 初始化Mysql
	initPostgres() // 初始化Postgres
	initRedis()    // 初始化Redis
}
