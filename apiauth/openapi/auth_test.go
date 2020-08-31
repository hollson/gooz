// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package openapi

import (
	"fmt"
	"net/url"
	"testing"
)

func TestURI(t *testing.T) {
	var (
		_url  *url.URL   // URL网址
		_vals url.Values // 参数列表
	)

	// 可对URL进行修改
	_url, _ = url.Parse("ftp://localhost:8080/a.html?name=张三&age=22&email=abc@126.com&Tel=13  9&住址=bj&sign=xyz")
	_url.Scheme = "https"
	_url.Host = "google.com"

	// 参数操作
	_vals = _url.Query()
	fmt.Println("参数原串：", _url.RawQuery)
	fmt.Println("参数字典：", _vals)

	_vals.Set("type", "ios")       // 替换/添加
	_vals.Add("tip", "yyy")        // 重复追加
	_vals.Add("tip", "xxx")        // 重复追加
	_vals.Del("age")               // 删除
	fmt.Println("读取参数：",_vals.Get("sign")) // 读取
	fmt.Println("排序编码：",_vals.Encode())
}
