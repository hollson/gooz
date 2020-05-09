package models

type Product struct {
	Id    int64  `xorm:"pk autoincr BIGINT(20)"`
	Name  string `xorm:"VARCHAR(50)"`
	Price string `xorm:"DECIMAL(10,2)"`
}
