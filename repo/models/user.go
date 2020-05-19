package models

import (
	"time"
)

type User struct {
	Id      int       `xorm:"not null pk autoincr INTEGER"`
	Name    string    `xorm:"VARCHAR(20)"`
	Created time.Time `xorm:"default now() DATETIME"`
	ClassId int       `xorm:"default 1 INTEGER"`
}
