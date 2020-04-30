//------------------------------------------------------------------
// @ Author: hollson <hollson@live.com>
// @ Date: 2019-12-01
// @ Version: 1.0.0
//------------------------------------------------------------------

package main

import (
	"github.com/hollson/deeplink/app"
	"runtime"
)

func main() {
	runtime.Gosched()
	app.Run()
}