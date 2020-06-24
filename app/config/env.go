// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// 运行环境
type Env string

// 运行环境常量
// 参考：https://en.wikipedia.org/wiki/Deployment_environment
const (
	Env_DEV   Env = "dev"   // 开发环境
	Env_TEST  Env = "test"  // 测试环境
	Env_Stage Env = "stage" // 验收环境
	Env_PROD  Env = "prod"  // 生产环境
)

// Env值映射为Gin的运行环境值
var Env2Gin = map[Env]string{
	Env_DEV:  gin.DebugMode,
	Env_TEST: gin.TestMode,
	Env_PROD: gin.ReleaseMode,
}

// 获取运行环境
// 获取系统环境变量GOENV，缺省值为
func GoEnv() Env {
	mode := os.Getenv("GOENV")
	mode = strings.ToLower(mode)
	if env := Env(mode); matchEnv(env) {
		return env
	}
	return Env_DEV
}

// 将运行环境映射为Gin的运行环境
func GinEnv() string {
	env:=GoEnv()
	switch env {
	case Env_DEV:
		return gin.DebugMode
	case Env_TEST:
		return gin.TestMode
	default:
		return gin.ReleaseMode
	}
}

func matchEnv(tar Env, envs ...Env) bool {
	for _, value := range envs {
		if value == tar {
			return true
		}
	}
	return false
}
