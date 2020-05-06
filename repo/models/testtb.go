package models

type Testtb struct {
	Id   int    `xorm:"INT(11)"`
	Name string `xorm:"VARCHAR(50)"`
}
