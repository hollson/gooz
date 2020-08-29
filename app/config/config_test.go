// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.


//https://www.kancloud.cn/liupengjie/go/743047

package config

import (
	"testing"

	"github.com/BurntSushi/toml"
)

func TestRedis(t *testing.T)  {
	var config struct{
		Sentinels sentinel `toml:"sentinel"`
	}

	if _, err := toml.Decode(TPL, &config); err != nil {
		t.Error(err)
	}
	t.Log(config)
}