package models

type Parts struct {
	PartNo      string `xorm:"not null pk VARCHAR(18)"`
	Description string `xorm:"VARCHAR(40)"`
	Cost        string `xorm:"not null DECIMAL(10,2)"`
	Price       string `xorm:"not null DECIMAL(10,2)"`
}
