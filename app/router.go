package app

import (
	"net/http"

	"github.com/EDDYCJY/go-gin-example/middleware/jwt"
	"github.com/gin-gonic/gin"
	"github.com/hollson/deeplink/app/export"
	"github.com/hollson/deeplink/service/help"
	"github.com/hollson/deeplink/service/user"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var (
	helper  *gin.RouterGroup // 帮助模块
	rgAsset *gin.RouterGroup // 静态资源
	rgUser  *gin.RouterGroup // 用户管理
	rgSys   *gin.RouterGroup // 系统管理
)

func Route() {
	// 帮助模块
	// Doc: http://localhost:8080/swagger/index.html
	helper = router.Group("/help")
	{
		helper.GET("/ping", help.PingHandler)
		helper.GET("/swg/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// 静态资源
	rgAsset = router.Group("/asset")
	{
		rgAsset.StaticFS("/favorite", http.Dir("/favorite.icon"))
		rgAsset.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
		rgAsset.StaticFS("/upload", http.Dir(export.GetExcelFullPath()))
		rgAsset.StaticFS("/download", http.Dir(export.GetExcelFullPath()))
	}

	// 用户模块
	rgUser = router.Group("/v1/user")
	rgUser.Use(jwt.JWT())
	{
		// 获取用户信息 curl http://127.0.0.1:8080/v1/user/detail?id=1
		rgUser.GET("/detail", user.GetUserHandler)
		rgUser.GET("/detail2", user.GetUserHandler)
	}

	// 其他模块...

	// 系统模块
	rgSys = router.Group("/v1/sys")
	rgSys.Use(jwt.JWT())
	{
		rgSys.GET("/profile", user.GetUserHandler)
		rgSys.GET("/logout", user.GetUserHandler)
	}
}
