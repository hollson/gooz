// -----------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.cn>
// @ Date: 2020-05-18
// @ Version: 1.0.0
//
// UserClass:自定义的Join联合对象
// -----------------------------------------------------------------------

package knit

import (
	"github.com/hollson/deeplink/repo/models"
)

// 自定义的Join联合对象
type ClassUser struct{
	models.User `xorm:"extends"`
	Name string
}
