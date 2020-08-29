// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/samuel/go-zookeeper/zk"
	"github.com/sirupsen/logrus"
)

var cli *zk.Conn

/**
*创建连接
 */
func TestConnect(t *testing.T) {
	// 集群
	var hosts = []string{"localhost:2181", "localhost:2182", "localhost:2183", "localhost:1234"}

	opt1 := zk.WithLogInfo(true)                // 是否开启日志
	opt2 := zk.WithHostProvider(&zk.DNSHostProvider{}) // 自定义Host程序，可更改集群主机读取方式

	conn, event, err := zk.Connect(hosts, 10*time.Second, opt1, opt2)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Println(<-event)
	cli = conn
}

func TestExists(t *testing.T) {
	// conn := Connect()
	cli.Exists()
	log.Println( cli.Server())
	defer cli.Close()
	fmt.Println(cli.Exists("/deeplink"))

}

//
// /**
// * 创建节点
// * 语法：create [-s] [-e] [-c] [-t ttl] path [data] [acl]
//  */
// func Create() {
// 	conn := Connect()
// 	defer conn.Close()
//
// 	var (
// 		path  = "/go"
// 		data  = []byte("hello golang")
// 		flags = 0                       // 0:永久；1:临时；2:自增序号；3:临时自增
// 		acls  = zk.WorldACL(zk.PermAll) // 访问权限模式
// 	)
// 	// conn.Create("/gotest", nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
// 	pth, err := conn.Create(path, data, int32(flags), acls)
// 	if err != nil {
// 		panic(err)
// 	}
// 	zk.ErrNodeExists
// 	fmt.Printf("创建成功: %v\n\n", pth)
// }
//
//
//
// /**
// * 修改节点
// * 语法：set [-s] [-v version] path data
// *      version参数用于CAS支持,即乐观锁
//  */
// func Set() {
// 	conn := Connect()
// 	defer conn.Close()
//
// 	stat, err := conn.Set("/go", []byte("hi golang"), -1)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("修改成功: %+v\n\n", stat)
// }
//
// /**
// * 删除节点 (不支持DeleteAll)
// * 语法：delete [-v version] path
//  */
// func Delete() {
// 	conn := Connect()
// 	defer conn.Close()
//
// 	err := conn.Delete("/go", -1)
// 	if err != nil {
// 		panic(err)
// 	} else {
// 		fmt.Println("删除成功")
// 	}
// }
//
// /**
// * 查询节点
//  */
// func Get() {
// 	conn := Connect()
// 	defer conn.Close()
//
// 	data, stat, err := conn.Get("/go")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("查询:%q,%+v\n\n", string(data), stat)
// }
//
// /**
// * 查询子节点名称(不递归)
//  */
// func Children() {
// 	conn := Connect()
// 	defer conn.Close()
//
// 	children, _, err := conn.Children("/go")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Printf("%+v \n", children)
// }
