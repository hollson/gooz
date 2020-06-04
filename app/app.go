package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hollson/deeplink/app/config"
	"github.com/hollson/deeplink/app/midware/stats"

	_ "github.com/hollson/deeplink/repo"
)

var router *gin.Engine

func Init() {
	fmt.Printf(WELCOME, config.App.Version, config.App.Env, config.App.Name)

	router = gin.Default()
	router.Use(gin.Logger())
	// router.Use(gin.Recovery())

	// router.Use(stats.ApiVisitHandler)
	Route()

	router.GET("/help/stats", stats.GetCurrentRunningStats)
	router.GET("/help/check", stats.GetCurrentRunningStats)

	// gin.SetMode(config.Env2Gin[config.App.Env])
}

// 启动程序
func Run() {
	server := &http.Server{
		Addr:           config.App.Port,
		Handler:        router,
		ReadTimeout:    time.Second * 30,
		WriteTimeout:   time.Second * 30,
		MaxHeaderBytes: 1 << 20, // 2M
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
