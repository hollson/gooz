// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package open

import (
	"time"
)

const (
	TERM = time.Second * 10
)

type claims struct {
	Mark string        // 客户简称
	Key  string        // 用户密钥
	Term time.Duration // 生效时长
	Free bool          // 是否免检
}

// Open Api 受众(编号:令牌)
var audience map[int64]*claims

// todo：配置文件（open_api）
func init() {
	audience = make(map[int64]*claims)
	audience[100000] = &claims{"测试1", "KhBQKxfMeElLPUNSQ66R0y9yW5kio4XM", TERM, true}
	audience[100011] = &claims{"客户A", "XhBQKxfMeE1LJUNSQ66R0y3yW5kio4AA", TERM, false}
	audience[100012] = &claims{"客户B", "YhBQKxfMeE2LKUNSQ66R0y4yW5kio4AB", TERM, false}
	audience[100013] = &claims{"客户C", "ZhBQKxfMeE3LLNMSQ66R0y5yW5kio4AC", TERM, false}
}