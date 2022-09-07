// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package open

import (
	"net/url"
)

// Authentication Authenticator 验证器
// verify

type OpenAPI struct {
	AppID     string `json:"app_id" form:"app_id"`         // 开发者账号(可选)
	AppKey    string `json:"app_key" form:"app_key"`       // 应用账号
	AppSecret string `json:"app_secret" form:"app_secret"` // 应用密码
}

// AppID：应用的唯一标识AppKey：公匙（相当于账号）AppSecret：私匙（相当于密码）
//
// token：令牌（过期失效）
// app_id, app_key, app_secret：我的身份证，银行卡号，银行密码。

// 安全API接口
type SecureApi interface {
	Verify() error
}

type openApi struct {
	uuid     string
	timeSpan int64
	sign     string
	token    string
	isFree   bool
	once     bool
}

func New(opts ...Option) *openApi {
	obj := &openApi{
		uuid:     "",
		timeSpan: 0,
		sign:     "",
		token:    "",
		isFree:   false,
		once:     false,
	}

	for _, val := range opts {
		val(obj)
	}

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	url.ParseQuery(u.token.RawQuery)

	return obj
}

func (p *openApi) Verify() error {

	return nil
}

// https://www.cnblogs.com/Sinte-Beuve/p/12093307.html
