package config

import (
	"github.com/gin-gonic/gin"
)

// Env值映射为Gin的运行环境值
var Env2Gin = map[Env]string{
	Env_DEV:  gin.DebugMode,
	Env_TEST: gin.TestMode,
	Env_PROD: gin.ReleaseMode,
}

func Init() {
	load()
}
