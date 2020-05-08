package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"
)

// MD5算法
func Md5V1(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	rt := hex.EncodeToString(h.Sum(nil))
	return rt
}

// 基于md5生成一个随机字符串
func TimeHash() string {
	t := time.Now()
	h := md5.New()
	io.WriteString(h, "crazyof.me")
	io.WriteString(h, t.String())
	passwd := fmt.Sprintf("%x", h.Sum(nil))
	return passwd
}

func Today() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
}

func RandIntn(n int) int {
	// 创建随机数种子  (默认将1970-1-1 00:00:00当做随机数种子)
	rand.Seed(time.Now().UnixNano()) // UnixNano()表示纳秒
	a := rand.Intn(n)                // 0-n之间的随机整数。 (包含0，不包含n。相当于对10求余)
	return a
}

func AtoInt64(num string) int64 {
	if nt64, ie := strconv.ParseInt(num, 10, 64); ie != nil {
		return 0
	} else {
		return nt64
	}
}

//hash := sha256.New()
//hash.Write([]byte(sign))
//md := hash.Sum(nil)
//mdStr := hex.EncodeToString(md)
