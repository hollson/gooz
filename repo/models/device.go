package models

type Device struct {
	Id     int64  `xorm:"pk autoincr BIGINT(20)"`
	Serial string `xorm:"not null VARCHAR(32)"`
}
