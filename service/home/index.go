// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package home

import (
	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context){
	c.String(200, `
       _              _ _      _
    __| |___ ___ _ __| (_)_ _ | |__
   / _  / -_) -_) '_ \ | | ' \| / /
   \____\___\___| .__/_|_|_||_|_\_\   
                |_|
`)
}