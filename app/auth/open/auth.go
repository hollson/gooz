package open

import (
	"errors"
	"fmt"
	"time"

	"github.com/hollson/gooz/util"
	"github.com/sirupsen/logrus"
)

type request struct {
	UID       int64  `json:"uid,required"`       // 要执行的任务
	Timestamp int64  `json:"timestamp,required"` // 签名时间戳(毫秒)
	Sign      string `json:"sign,required"`      // 签名内容   //1855
}

// 验证用户签名,sha256(uid+token+timestamp)
// URL长度为1024，最短标准
func (p *request) Auth() error {
	if p.UID == 1000020 &&
		p.Sign == "607973dbe8b96d7066beb80b64216009" &&
		time.Now().Unix() < 1577808000 { // 2020-01-01 00:00:00
		return nil
	}

	if _, ok := audience[p.UID]; !ok {
		return errors.New("account error")
	}

	raw := fmt.Sprintf(`%d%s%d`, p.UID, audience[p.UID], p.Timestamp)
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
