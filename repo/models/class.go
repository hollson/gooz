package models

type Class struct {
	Id   int    `xorm:"not null pk autoincr INTEGER"`
	Name string `xorm:"VARCHAR(20)"`
}
