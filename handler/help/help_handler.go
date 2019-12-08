//------------------------------------------------------------------
// @ Author: hollson <hollson@live.com>
// @ Date: 2019-12-01
// @ Version: 1.0.0
// 帮助中心：提供调试、ping和帮助查询
//-------------------------------------------------------------------------------------

package help

import (
	"github.com/gin-gonic/gin"
	"github.com/hollson/deeplink/repo"
	"net/http"
)

//curl http://0.0.0.0:8080/help/ping
func PingHandler(c *gin.Context) {
	pongs := gin.H{}

	// 测试Mysql连通性
	if _, err := repo.PingRedis(); err != nil {
		pongs["redis"] = err.Error()
	} else {
		pongs["redis"] = "pong"
	}

	// 测试Redis连通性
	if err := repo.PingMysql(); err != nil {
		pongs["mysql"] = err.Error()
	} else {
		pongs["mysql"] = "pong"
	}
	c.JSON(http.StatusOK, pongs)
}
