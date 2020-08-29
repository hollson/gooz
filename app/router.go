package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hollson/gooz/app/auth/jwt"
	"github.com/hollson/gooz/app/export"
	"github.com/hollson/gooz/app/midware/stats"
	"github.com/hollson/gooz/service/account"
	"github.com/hollson/gooz/service/article"
	 "github.com/hollson/gooz/service/home"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var (
	rgAsset *gin.RouterGroup // 静态资源
	rgAtc   *gin.RouterGroup // 文章管理
	rgSys   *gin.RouterGroup // 系统管理
	helper  *gin.RouterGroup // 帮助模块
)

func Route() {
	// main
	main := router.Group("/")
	{
		main.GET("",home.IndexHandler)
	}

	// 游客身份
	guest := router.Group("/v1/account")
	{
		guest.POST("/register", article.GetArticleDetailHandler) // 注册
		guest.POST("/login", account.LoginHandler)               // 登录
		guest.POST("/forget", article.GetArticleDetailHandler)   // 忘记
	}

	// 账号管理
	acc := router.Group("/v1/account")
	acc.Use(jwt.Auth())
	{
		// 获取用户信息 curl http://127.0.0.1:8080/v1/atc/detail?id=1
		acc.GET("/profile", article.GetArticleDetailHandler) // 个人信息
		acc.GET("/modpwd", article.GetArticleDetailHandler)  // 修改密码
		acc.GET("/logout", article.GetArticleDetailHandler)  // 登出
	}

	// 主页
	home := router.Group("/v1/account")
	home.Use(jwt.Auth())
	{
		// 获取用户信息 curl http://127.0.0.1:8080/v1/atc/detail?id=1
		home.GET("/home/index", article.GetArticleDetailHandler)
		home.GET("/home/detail", article.GetArticleDetailHandler)
	}

	// 文章模块
	rgAtc = router.Group("/v1/atc")
	rgAtc.Use(jwt.Auth())
	{
		// 获取用户信息 curl http://127.0.0.1:8080/v1/atc/detail?id=1
		rgAtc.GET("/detail", article.GetArticleDetailHandler)
		rgAtc.GET("/list", article.GetArticleDetailHandler)
	}

	// 系统模块
	rgSys = router.Group("/v1/sys")
	rgSys.Use(jwt.Auth())
	{
		rgSys.GET("/profile", article.GetArticleDetailHandler)
		rgSys.GET("/logout", article.GetArticleDetailHandler)
	}

	// 静态资源
	rgAsset = router.Group("/asset")
	{
		rgAsset.StaticFS("/favorite", http.Dir("/favorite.icon"))
		rgAsset.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
		rgAsset.StaticFS("/upload", http.Dir(export.GetExcelFullPath()))
		rgAsset.StaticFS("/download", http.Dir(export.GetExcelFullPath()))
	}

	// 帮助模块
	helper = router.Group("/help")
	{
		// 自动文档: http://localhost:8080/help/doc/index.html
		helper.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		// 系统监控: http://localhost:8080/help/stats
		helper.GET("/stats", stats.GetCurrentRunningStats)
	}
}
