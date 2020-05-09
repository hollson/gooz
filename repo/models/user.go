package models

type User struct {
	Id   int64  `xorm:"pk autoincr BIGINT(11)"`
	Name string `xorm:"VARCHAR(50)"`
}
