//------------------------------------------------------------------
// @ Author: hollson <hollson@live.com>
// @ Date: 2019-12-01
// @ Version: 1.0.0
//------------------------------------------------------------------

package app

import "github.com/hollson/deeplink/handler/help"

func Route() {
	// 测试连通性
	router.GET("/help/ping", help.PingHandler)
}
