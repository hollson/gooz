package repo

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/hollson/deeplink/app/config"
	"github.com/sirupsen/logrus"
)

func initRedis() {
	for k, v := range config.Redis {
		rds, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", v.Host, v.Port))
		if err != nil {
			logrus.Errorf(" ❌  Redis-%s 拨号失败,%s", k, err.Error())
			continue
		}

		if _, err := redis.String(rds.Do("ping")); err != nil {
			logrus.Errorf(" ❌  Redis-%s [%s:%d] 连接失败,%s", k, v.Host, v.Port, err.Error())
			continue
		}

		logrus.Infof(" ✅  Redis-%s [%s:%d] 连接成功", k, v.Host, v.Port)
		Rds = append(Rds, rds)
	}

	test()
}

// 测试
func test() {
	defer func() {
		if err := recover(); err != nil {
			logrus.Errorln(err)
		}
	}()
	SetValue("hi", "hello world")
	rt, err := GetValue("hi")
	fmt.Println("Redis测试", rt, err)
}

func SetValue(key, val string) {
	// 24+小时过期 todo
	Rds[0].Do("SET", key, val)
}

func GetValue(key string) (string, error) {
	return redis.String(Rds[0].Do("GET", key))
}
