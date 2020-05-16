package models

type Product struct {
	Id   int64  `xorm:"pk BIGINT"`
	Name string `xorm:"VARCHAR"`
}
