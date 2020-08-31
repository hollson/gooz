package app

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hollson/gooz/app/config"
	"github.com/hollson/gooz/app/midware/stats"
	"github.com/hollson/goox/color"
	"github.com/sirupsen/logrus"

	_ "github.com/hollson/gooz/repo"
)

const (
	// The welcome screen
	// usage：
	// 		fmt.Printf(WELCOME,app.version,app.env,app.name)
	// See http://www.network-science.de/ascii/ for more .
	WELCOME = `
Deeplink is a internal applications <XX科技.保留版权>	
   ____  ____  ____  ____
  / __ \/ __ \/ __ \/_  /
 / /_/ / /_/ / /_/ / / /_
 \__, /\____/\____/ /___/
/____/    		 %s(%s)

Usage:
	%s <command> [arguments] 

Use "Gooz help <command>" for more information about a command.
For more please email hollson@qq.com

`
)

var (
	router *gin.Engine
)

func Init() {
	fmt.Printf(color.Blue( WELCOME), config.App.Version, config.App.Env, config.App.Name)
	router = gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(stats.ApiVisitHandler)
	gin.SetMode(config.GinEnv())
	// GinDump()
	Route()
}

// Gin日志重定向
func GinDump() {
	logfile, err := os.OpenFile("./logs/gin.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	logrus.SetOutput(logfile)

	gin.DisableConsoleColor()                              // 不需要颜色
	gin.DefaultWriter = io.MultiWriter(os.Stdout, logfile) // os.Stdout
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

	fmt.Println(" 🚗 服务已启动", config.App.Port)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
