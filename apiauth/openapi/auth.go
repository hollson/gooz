// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package openapi

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"openapi/bloom"
)

// 开放API授权认证规则说明：
// 安全层次一：验证身份
// 只验证时间戳和用户标识，即确保的授权用户的操作
//
// 参考阿里开放API文档：https://help.aliyun.com/document_detail/29475.html
// appSecert各方共同持有，不参与通讯  identity

var (
	bloomer *bloom.BloomFilter
)

type Auth struct {
	UID       string     `json:"uid,required"`       // 用户标识
	Timestamp int64      `json:"timestamp,required"` // 时间戳(毫秒)
	Nonce     int64      `json:"nonce"`              // 随机数
	Sign      string     `json:"sign,required"`      // 签名
	Values    url.Values `json:"params"`             // 参数列表
}

// API安全三步骤：
// 1.身份验证
// 2.内容篡改
// 3.重放攻击
type validator struct {
	base     bool          // 身份验证（简单验证）
	fake     bool          // 防篡改（递进可选）
	replay   bool          // 防重放（递进可选）
	lifetime time.Duration // 有效期(默认10秒)
}

func New(opts ...Option) *validator {
	v := &validator{
		base:     true,
		fake:     false,
		replay:   false,
		lifetime: time.Second * 10,
	}
	for _, f := range opts {
		f(v)
	}
	if v.replay {
		bloomer = bloom.NewBloomFilter()
	}
	return v
}

func (p *validator) Verify(auth *Auth) error {
	// 随机数
	// 时间戳
	// 用户
	// 签名
	//
	return nil
}

// 安全层次一：放参数被篡改
// 覆盖第一层次验证，即校验所有参数项，验证策略如下：
//
func (p *validator) check(auth *Auth) error {
	if p.lifetime < time.Second*3 {
		return errors.New("token lifetime can not less than 3 seconds")
	}
	// if p.replay&&p.
	return nil
}

// 安全层次一：基础验证
// 只验证时间戳和用户标识，即确保的授权用户的操作
func (p *validator) baseVerify(ident Identity,auth *Auth) error {
	// if p.lifetime < time.Second*3 {
	// 	return errors.New("token lifetime can not less than 3 seconds")
	// }

// tar:=fmt.Sprintf("%s%s",ident.AppKey,ident.AppSecret)

	if p.fake {
		return p.fakeVerify(auth)
	}

	// ???
	if p.replay {
		mark := fmt.Sprintf("%s%d%d", auth.UID, auth.Timestamp, auth.Nonce)
		bloomer.Add(mark)
	}
	return nil
}

// 安全层次一：放参数被篡改
// 覆盖第一层次验证，即校验所有参数项，验证策略如下：
//
func (p *validator) fakeVerify(auth *Auth) error {
	mark := fmt.Sprintf("%s%d%d", auth.UID, auth.Timestamp, auth.Nonce)
	if bloomer.Has(mark) {
		return errors.New("repeat submission error")
	}
	return nil
}

// 重复提交（使用布隆过滤器）
func (p *validator) replayVerify(auth *Auth) error {
	mark := fmt.Sprintf("%s%d%d", auth.UID, auth.Timestamp, auth.Nonce)
	if bloomer.Has(mark) {
		return errors.New("repeat submission error")
	}
	return nil
}

//
// // 验证用户签名,sha256(uid+token+timestamp)
// // URL长度为1024，最短标准
// func (p *Needs) Auth() error {
// 	if p.UID == 1000020 &&
// 		p.Sign == "607973dbe8b96d7066beb80b64216009" &&
// 		time.Now().Unix() < 1577808000 { // 2020-01-01 00:00:00
// 		return nil
// 	}
//
// 	if _, ok := audience[p.UID]; !ok {
// 		return errors.New("account error")
// 	}
//
// 	raw := fmt.Sprintf(`%d%s%d`, p.UID, audience[p.UID], p.Timestamp)
// 	mdStr := util.Md5V1(raw)
//
// 	if mdStr != p.Sign {
// 		logrus.Errorln(fmt.Sprintf("sign invalid , %v", p))
// 		return errors.New("sign invalid")
// 	}
//
// 	// N秒签名过期  todo 配置
// 	if time.Since(time.Unix(p.Timestamp, 0)) > time.Second*30 { // 30秒
// 		return errors.New("sign expired error")
// 	}
//
// 	return nil
// }
