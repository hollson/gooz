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
)

var (
	_home  *gin.RouterGroup // 通行游客
	_asset *gin.RouterGroup // 静态资源
	_prof  *gin.RouterGroup // 个人信息
	_sys   *gin.RouterGroup // 系统管理
	_help  *gin.RouterGroup // 帮助模块
)

func Route() {
	// 通行游客
	_home = router.Group("/")
	{
		_home.GET("", home.IndexHandler)
		_home.POST("/register", article.GetArticleDetailHandler) // 注册
		_home.POST("/login", account.LoginHandler)               // 登录
		_home.POST("/forget", article.GetArticleDetailHandler)   // 忘记密码
	}

	// 账号管理
	_prof = router.Group("/v1/profile")
	_prof.Use(jwt.Auth())
	{
		_prof.GET("/info", article.GetArticleDetailHandler)      // 个人信息
		_prof.GET("/mod/tel", article.GetArticleDetailHandler)   // 修改手机
		_prof.GET("/mod/email", article.GetArticleDetailHandler) // 修改邮箱
		_prof.GET("/mod/pwd", article.GetArticleDetailHandler)   // 修改密码
		_prof.GET("/logout", article.GetArticleDetailHandler)    // 登出账号
	}

	// 系统模块
	_sys = router.Group("/v1/sys")
	_sys.Use(jwt.Auth())
	{
		_sys.GET("/limit", article.GetArticleDetailHandler) // 权限管理
		_sys.GET("/set", article.GetArticleDetailHandler)   // 系统设置
		_sys.GET("/log", article.GetArticleDetailHandler)   // 日志管理
	}

	// 静态资源
	_asset = router.Group("/asset")
	{
		_asset.StaticFS("/css", http.Dir(export.GetExcelFullPath()))    // 样式
		_asset.StaticFS("/js", http.Dir(export.GetExcelFullPath()))     // 脚本
		_asset.StaticFS("/html", http.Dir(export.GetExcelFullPath()))   // 视图
		_asset.StaticFS("/upload", http.Dir(export.GetExcelFullPath())) // 上传
	}

	// 帮助模块(超级权限)
	_help = router.Group("/help")
	_help.Use(jwt.Auth())
	{
		_help.GET("/ping", stats.GetCurrentRunningStats)  // 检测系统
		_help.GET("/stats", stats.GetCurrentRunningStats) // 系统监控
	}
}
