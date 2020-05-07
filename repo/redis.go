package repo

import (
	"github.com/garyburd/redigo/redis"
	"github.com/sirupsen/logrus"
)

//todo
const redis_addr = "0.0.0.0:6379"

var rds redis.Conn

func init() {
	var err error
	rds, err = redis.Dial("tcp", redis_addr)

	if err != nil {
		logrus.Errorln(" ❌ Redis连接失败:", err)
	}else {
		if _,err:=PingRedis();err!=nil{
			logrus.Errorln(" ❌ Ping Redis失败:", err)
		}else {
			logrus.Infoln(" ✅ Postgres数据库连接成功 !!!")
		}
	}
}


func PingRedis() (string, error) {
	return redis.String(rds.Do("ping"))
}

func SetValue(key, val string) {
	// 24+小时过期 todo
	rds.Do("SET", key, val)
}

func GetValue(key string) (string, error) {
	return redis.String(rds.Do("GET", key))
}
