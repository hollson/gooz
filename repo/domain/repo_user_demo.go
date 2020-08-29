// -----------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.cn>
// @ Date: 2020-05-18
// @ Version: 1.0.0
//
// XORM操作示例：
// https://blog.csdn.net/fwhezfwhez/article/details/79291492
// https://github.com/go-xorm/xorm/blob/master/README_CN.md
// -----------------------------------------------------------------------

package domain

// import (
// 	"fmt"
// 	"strconv"
// 	"time"
//
// 	"github.com/hollson/gooz/repo/knit"
// 	"github.com/hollson/gooz/repo/models"
// 	_ "github.com/lib/pq"
// 	"xorm.io/xorm"
// )
//
// // 插入操作"
// func Insert() {
// 	// 1. 对象(批量)插入
// 	if false {
// 		user := &models.User{Name: "insert01"}
// 		db.Insert(user) // 可批量插入
// 	}
//
// 	// 2. 明文SQL插入(单次事务)
// 	if false {
// 		sql := "insert into public.user(name) values(?)"
// 		db.Exec(sql, "insert02")
// 	}
//
// 	// 3. 预编译SQL插入(单次事务)，Prepare详见：https://www.jianshu.com/p/ee0d2e7bef54
// 	if false {
// 		sql := "insert into public.user(name) values(?)"
// 		db.Prepare().SQL(sql, "insert03").Exec()
// 	}
// }
//
// func Select() {
// 	// 1. 单条查询(对象），由对象决定表名
// 	if false {
// 		user := new(User)
// 		ok, err := db.Where("id=?", 5).And("name=?", "ft7").Get(user)
// 		fmt.Println(ok, err, user)
// 	}
//
// 	// 2. 多条查询(限定字段)
// 	if false {
// 		var users []User
// 		err := db.Where("name = ?", "tom").
// 			And("id > 10").		// 条件
// 			Limit(10, 0).  	// 分页
// 			Cols("id,name"). 	// 部分列
// 			// AllCols().  				// 全部列
// 			Find(&users)
// 		fmt.Println(users, err)
// 	}
//
// 	// 3. 联合查询
// 	if false {
// 		users := new([]knit.ClassUser)
// 		// db.Join("INNER", "class", "user.class_id=class.id").Find(users)  //表B
// 		db.Table("users").Alias("a").
// 			Join("inner", "class", "a.class_id=b.id").Alias("b").
// 			Find(users)
// 		fmt.Println(users)
// 	}
//
// 	// 4. SQL查询
// 	if false {
// 		users := new([]knit.ClassUser)
// 		db.SQL("select u.id,u.name,c.name from public.user as u left join public.class as c on u.class_id=c.id").
// 			Find(users)
// 		fmt.Println(users)
// 	}
//
// 	// 5 链式查找
// 	if false {
// 		// 值得一提的是，支持查找某行的某个字段，不过一般在sql语句中就可以完成过滤，如果sql语句过于复杂，可以链式查找过滤
// 		db.SQL("select * from public.user where id<?", 10).Query("")
// 		rt, err := db.SQL("select * from public.user").Query() // .Results[0]["id"]
// 		fmt.Println(rt, err)
// 	}
//
// 	// 6. 游标遍历
// }
//
// // 是否存在
// func Exists() {
// 	// has, err := testEngine.Exist(new(RecordExist))
// 	// // SELECT * FROM record_exist LIMIT 1
// 	//
// 	// has, err = testEngine.Exist(&RecordExist{
// 	// 	Name: "test1",
// 	// })
// 	// // SELECT * FROM record_exist WHERE name = ? LIMIT 1
// 	//
// 	// has, err = testEngine.Where("name = ?", "test1").Exist(&RecordExist{})
// 	// // SELECT * FROM record_exist WHERE name = ? LIMIT 1
// 	//
// 	// has, err = testEngine.SQL("select * from record_exist where name = ?", "test1").Exist()
// 	// // select * from record_exist where name = ?
// 	//
// 	// has, err = testEngine.Table("record_exist").Exist()
// 	// // SELECT * FROM record_exist LIMIT 1
// 	//
// 	// has, err = testEngine.Table("record_exist").Where("name = ?", "test1").Exist()
// 	// // SELECT * FROM record_exist WHERE name = ? LIMIT 1
// }
//
// // 调用函数
//
// func main() {
//
// 	// 6.执行更新
// 	// 6.1 ORM方式: 只有非0值的属性会被更新，user的id和created都是默认零值，不被处理
// 	if false {
// 		user := new(User)
// 		user.Name = "ftx"
// 		// [xorm] [info]  2018/02/08 12:04:01.330624 [SQL] UPDATE "user" SET "name" = $1 WHERE "id"=$2 []interface {}{"ftx", 4}
// 		db.Id(4).Update(user)
// 	}
// 	// 6.2 SQL方式略,和insert类似
//
// 	// 7.事务
// 	// 7.1简单事务
// 	if false {
// 		session := db.NewSession()
// 		defer session.Close()
//
// 		session.Begin()
// 		// 业务:新添加学生，并且创建新的班级，如果班级因为主键冲突创建失败，则整个事务回滚
// 		_, err = session.SQL("insert into public.user(name,class_id) values('ft13',2)").Execute()
// 		// 表中已经有id=3的班级了
// 		_, err = session.SQL("insert into public.class(id,name) values(3,'高中3班')").Execute()
// 		if err != nil {
// 			session.Rollback()
// 		}
// 		session.Commit()
//
// 	}
// 	// 7.2嵌套事务
// 	if false {
// 		session := db.NewSession()
// 		defer session.Close()
// 		session.Begin()
// 		_, err := session.Exec("insert into public.user(name,class_id) values('ft23',2)")
// 		if err != nil {
// 			session.Rollback()
// 		}
// 		_, err = session.Exec("insert into public.user(id,name,class_id) values(1,'ft24',2)")
// 		if err != nil {
// 			session.Rollback()
// 		}
//
// 		tx, _ := session.BeginTrans()
// 		_, err = tx.Session().Exec("insert into public.user(name,class_id) values('ft25',2)")
// 		if err != nil {
// 			tx.RollbackTrans()
// 		}
// 		tx.CommitTrans()
// 		session.Commit()
//
// 	}
// 	// 8.缓存:使用Raw方式修改以后，需要清理缓存
// 	if true {
//
// 		// 建立500条数据
// 		session := db.NewSession()
// 		defer session.Close()
// 		if false {
// 			session.Begin()
// 			for i := 30; i < 530; i++ {
// 				value := "ft" + strconv.Itoa(i)
// 				_, err = session.Exec("insert into public.user(name) values(?)", value)
// 				if err != nil {
// 					session.Rollback()
// 				}
// 			}
// 			session.Commit()
// 		}
//
// 		// 查询前531条数据，并随意输出其中一条
// 		users := make([]User, 10)
// 		db.SQL("select * from public.user where id<531 order by id").Find(&users)
// 		fmt.Println("读第一遍:", "id:", users[50].Id, "name:", users[50].Name)
//
// 		db.SQL("select * from public.user where id<531 order by id").Find(&users)
// 		fmt.Println("读第二遍:", "id:", users[50].Id, "name:", users[50].Name)
//
// 		var step int = 1
// 		stepString := users[50].Name + strconv.Itoa(step)
// 		session.Exec("update public.user set name=? where id =45", stepString)
//
// 		// 清理缓存
// 		db.ClearCache(new(User))
//
// 		time.Sleep(5 * time.Second)
//
// 		session.SQL("select * from public.user where id<531 order by id").Find(&users)
// 		fmt.Println("读第三遍:", "id:", users[50].Id, "name:", users[50].Name)
//
// 		// 虽然很不好意思，但是就算开启了缓存数据也是脏了
//
// 	}
// 	// 9.读写分离
// 	if false {
// 		// 假设有多台服务器用来响应客户的读请求
// 		var dbGroup *xorm.EngineGroup
// 		conns := []string{
// 			"postgres://postgres:123@localhost:5432/test?sslmode=disable",
// 			"postgres://postgres:123@localhost:5432/test?sslmode=disable",
// 			"postgres://postgres:123@localhost:5432/test?sslmode=disable",
// 			"postgres://postgres:123@localhost:5432/test?sslmode=disable",
// 		}
//
// 		// 负载均衡策略:(特性自行百度)
// 		// 1.xorm.RandomPolicy()随机访问负载均衡,
// 		// 2.xorm.WeightRandomPolicy([]int{2, 3,4})权重随机负载均衡
// 		// 3.xorm.RoundRobinPolicy() 轮询访问负载均衡
// 		// 4.xorm.WeightRoundRobinPolicy([]int{2, 3,4}) 权重轮训负载均衡
// 		// 5.xorm.LeastConnPolicy()最小连接数负载均衡
// 		dbGroup, err = xorm.NewEngineGroup("postgres", conns, xorm.RoundRobinPolicy())
// 		// dbGroup使用方法和db一致
// 		// 简单查询
// 		dbGroup.SQL("inser into public.users(name) values('ft2000')").Execute()
// 		dbGroup.Exec("inser into public.users(name) values('ft2001')")
//
// 		// 事务查询
// 		session := dbGroup.NewSession()
// 		defer session.Close()
// 		session.Begin()
// 		_, err = session.Exec("inser into public.users(name) values('ft2001')")
// 		if err != nil {
// 			session.Rollback()
// 		}
// 		session.Commit()
// 	}
// }
//
// // 注意:
// // 1.postgresql好像不会默认按id增长排序，所以书写sql语句要提前写好order by id ，楼主没怎么写，咳咳
// // 2. [5.4] postgresql建表会建在public策略的table里，所以查询语句表明写的是public.xxxx,这也造成了连表orm查询会发生前缀报错,比如变成了"SELECT * FROM "public"."user" INNER JOIN "class" ON user.class_id=class.id 这和内部的split有关，
// // 3.[8.]带的缓存好像很容易失效,在创建500个数据后，经过查查改查的操作，查询到的结果是一样的始终是一样的，本来改值之后应该最后一遍查会变化，然而并没有,缓存功能即使清理了缓存，还是会读到脏的
