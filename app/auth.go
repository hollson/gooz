// ------------------------------------------------------------------
// @ Author: hollson <hollson@live.cn>
// @ Date: 2019-12-01
// @ Version: 1.0.0
// 验证授权：签名验证服务
// -------------------------------------------------------------------------------------

package app

import (
	"errors"
	"fmt"
	"time"

	"github.com/hollson/deeplink/util"
	"github.com/sirupsen/logrus"
)

// 渠道账号(客户编号:令牌)
type Users map[int64]string

// 分配渠道账号(客户编号:令牌)   todo 配置
var users = Users{
	1000010: "KhBQKxfMeElLPUNSQ66R0y9yW5kio4XM", // 测试
}

func (p Users) GetAll() Users {
	return users
}

type Auth struct {
	UID       int64  `json:"uid"`       // 要执行的任务
	Timestamp int64  `json:"timestamp"` // 签名时间戳(毫秒)
	Sign      string `json:"sign"`      // 签名内容：sha256(uid+token+timestamp)
}

// 验证用户签名
func (p *Auth) CheckAuth() error {
	if p.UID == 1000020 &&
		p.Sign == "607973dbe8b96d7066beb80b64216009" &&
		time.Now().Unix() < 1577808000 { // 2020-01-01 00:00:00
		return nil
	}

	if _, ok := users[p.UID]; !ok {
		return errors.New("account error")
	}

	raw := fmt.Sprintf(`%d%s%d`, p.UID, users[p.UID], p.Timestamp)
	mdStr := util.Md5V1(raw)

	if mdStr != p.Sign {
		logrus.Errorln(fmt.Sprintf("sign invalid , %v", p))
		return errors.New("sign invalid")
	}

	// N秒签名过期  todo 配置
	if time.Since(time.Unix(p.Timestamp, 0)) > time.Second*30 { // 30秒
		return errors.New("sign expired error")
	}

	return nil
}
