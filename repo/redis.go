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
		logrus.Fatal("Redis连接失败:", err)
		panic(err)
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

func FlushDb(index int) {
	//rds.Do("SET", key, val)
}
