// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package openapi

type Passport interface {

}

// App客户端身份认证信息
type Identity struct {
	AppKey    string // 标识
	AppSecret string // 密钥
	Free      bool   // 免检
	Mark      string // 备注
}

var Users map[string]Identity

// todo：配置文件（open_api）
func init() {
	us := []Identity{
		{"100011", "MHBQKX90EELLPUNUI66R0Y9YW5KIOXXX", false, "编辑器"},
		{"100012", "XHBQKXFMEE1LJUNSQ66R0Y3YW5KIO4AA", false, "安卓端"},
		{"100013", "YHBQKX8IEE2LKUNSQ16R0Y4YW5KIO4AB", false, "IOS端"},
		{"100014", "KHBQKXFMEELLPUNSQ56R0Y9YW5KIO4XM", false, "WEB端"},
		{"999999", "ZHBQKXFMEE3LLNMSQ86R0Y5YW5KIO4AC", true, "测试"},
	}

	Users = make(map[string]Identity,len(us))
	for _, val := range us {
		Users[val.AppKey]=val
	}
}
