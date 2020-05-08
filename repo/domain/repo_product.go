//-------------------------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.com>
// @ Date: 2020-05-07
// @ Version: 1.0.0
//
// Here's the code description...
//-------------------------------------------------------------------------------------

// 产品数据仓储
package domain

import (
	"errors"
	"github.com/hollson/deeplink/repo"
	"github.com/hollson/deeplink/repo/models"
	"github.com/sirupsen/logrus"
)

type IProduct interface {
	GetProduct(id int64) (*models.Product, error)
	GetProductByPage(offset, limit int) ([]models.Product, error)
}

type  ProductRepo struct {}

func (p *ProductRepo) GetProduct(id int64) (*models.Product, error) {
	u:=models.Product{Id:id}
	if _,err:=repo.Eg.Get(&u);err!=nil{
		return nil,errors.New("")
	}else {
		return &u,nil
	}
}

func (p *ProductRepo) GetProductByPage(offset, limit int) ([]models.Product, error) {
	var ps []models.Product

	if err:=repo.Eg.SQL("select * from product limit 10;").Find(ps);err!=nil{
		logrus.Errorln("GetProductByPage Err：",err)
		return nil,err
	}
	return ps,nil
}
