// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package openapi

import (
	"time"
)

type Option func(opt *validator)

// 防止参数篡改
func WithNoFake(sure bool) Option {
	return func(opt *validator) {
		opt.fake = sure
	}
}

// 方式重放攻击
func WithNoReplay(sure bool) Option {
	return func(opt *validator) {
		opt.replay = sure
	}
}

// Token有效期(默认十秒，默认10秒，不能小于3秒)
func WithLifetime(lifetime time.Duration) Option {
	return func(opt *validator) {
		opt.lifetime = lifetime
	}
}
