#!/usr/bin/env bash
#--------------------------------------------------------------------------
# @ Copyright (C) XX科技 - 保留版权 .
# @ Author: hollson <hollson@live.com>
# @ Date: 2019-12-01
# @ Describe：使用golang-xorm框架的代码生成器，生成数据库实体对象
#--------------------------------------------------------------------------

templates=$GOPATH/src/github.com/go-xorm/cmd/xorm/templates/goxorm/

sudo rm -rf ./models/*
sudo xorm reverse mysql root:"123456"@"(127.0.1:3306)"/dbtest?charset=utf8 $templates ./models