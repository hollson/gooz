package models

type Product struct {
	Name  string `xorm:"not null VARCHAR(50)"`
	Id    int64  `xorm:"default nextval('"Product_id_seq"'::regclass) autoincr BIGINT"`
	Price string `xorm:"default 0 NUMERIC"`
}
