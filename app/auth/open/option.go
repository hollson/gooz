// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package open


type Option func(opt *openApi)

func WithUuid(uuid string) Option {
	return func(opt *openApi) {
		opt.uuid = uuid
	}
}

func WithTimeSpan(timespan int64) Option {
	return func(opt *openApi) {
		opt.timeSpan = timespan
	}
}

func WithSign(sign string) Option {
	return func(opt *openApi) {
		opt.sign = sign
	}
}

func WithToken(token string) Option {
	return func(opt *openApi) {
		opt.token = token
	}
}

func WithFree(free bool) Option {
	return func(opt *openApi) {
		opt.isFree = free
	}
}

func WithOnce(once bool) Option {
	return func(opt *openApi) {
		opt.once = once
	}
}

// type OpenApi interface {
// 	// UUID string
// 	GetUUID() string    // 用户标识
// 	GetTimeSpan() int64 // 签名时间戳
// 	GetSign() string    // 签名内容
// 	GetToken() []byte   // 客户令牌
// 	IsFree() bool // 是否免检
//
// 	Verify() error
// 	// Once  //单次有效
// 	// Term time.Duration // 生效时长
// }