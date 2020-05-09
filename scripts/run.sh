#!/bin/bash

# 停止当前运行的服务
pkill tmp_appname

# 重新启动服务
./tmp_appname -d=true

# 查看进程
ps aux|grep tmp_appname

echo " ✈️ 服务已开启..."