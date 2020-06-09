// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package account

import (
	"errors"
)

type UserInfo struct {
	Id     int64
	Name   string
	Tel    string
	Email  string
	Sex    string
	Role   int64
	Limits []int64
}

// 模拟数据
func UserDemo(req *ReqLogin) (*UserInfo, error) {
	if req.Account != "admin" && req.Password != "123456" {
		return nil, errors.New("username or password invalid")
	}
	return &UserInfo{
		Id:     1001,
		Name:   "hollson",
		Role:   28,
		Tel:    "18201188888",
		Email:  "hollson@qq.com",
		Sex:    "M",
		Limits: []int64{1, 3, 12, 24, 16},
	}, nil
}
