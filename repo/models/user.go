package models

type User struct {
	Id   int64  `xorm:"pk BIGINT"`
	Name string `xorm:"not null VARCHAR(20)"`
}
