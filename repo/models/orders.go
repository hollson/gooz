package models

import (
	"time"
)

type Orders struct {
	Id      int       `xorm:"not null pk INTEGER"`
	Created time.Time `xorm:"DATE"`
	Content string    `xorm:"VARCHAR(255)"`
}
