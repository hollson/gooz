// ------------------------------------------------------------------
// @ Author: hollson <hollson@live.cn>
// @ Date: 2019-12-01
// @ Version: 1.0.0
// ------------------------------------------------------------------

package main

import (
	"runtime"

	"github.com/hollson/deeplink/app"
)

func main() {
	runtime.Gosched()
	app.Run()
}
