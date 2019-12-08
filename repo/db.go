package repo

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/hollson/deeplink/config"
	"github.com/sirupsen/logrus"
)

//定义orm引擎
var DB *xorm.Engine

//创建orm引擎
func init() {
	var err error
	DB, err = xorm.NewEngine("mysql", config.DbSource)
	if err != nil {
		logrus.Fatal("数据库连接失败:", err)
		panic(err)
	}
	DB.ShowSQL(true)
}

// 测试连通性
func PingMysql() error {
	return DB.Ping()
}
