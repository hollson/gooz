// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package open

import (
	"net/url"
)



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
