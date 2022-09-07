// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/BurntSushi/toml"
)

// 全局变量
var (
	App     *app
	System  *system
	Service map[string]*service
	Log     *log
	Repo    map[string]*repo
	Redis   *redis
	Kafka   *kafka
)

// 全局配置
type configure struct {
	App     *app                `toml:"app"`
	System  *system             `toml:"system"`
	Service map[string]*service `toml:"service"`
	Log     *log                `toml:"log"`
	Repo    map[string]*repo    `toml:"repo"`
	Redis   *redis              `toml:"redis"`
	Kafka   *kafka              `toml:"kafka"`
}

// Load toml config, Usage:
//  func main() {
//  	if err := config.Load(os.Args[1]); err != nil {
//  		panic(err)
//  	}
//  	fmt.Println(config.App)
//  	fmt.Println(config.System)
//  	fmt.Println(config.Service)
//  	fmt.Println(config.Repo)
//  	fmt.Println(config.Log)
//  	fmt.Println(config.Redis)
//  	fmt.Println(config.Kafka)
//  }
func Load(path string) error {
	var c configure
	if _, err := toml.DecodeFile(path, &c); err != nil {
		return err
	}
	App = c.App
	System = c.System
	Service = c.Service
	Log = c.Log
	Repo = c.Repo
	Redis = c.Redis
	Kafka = c.Kafka
	return nil
}
