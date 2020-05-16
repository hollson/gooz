#!/bin/bash

# 停止当前运行的服务
pkill tmp_appname

# 重新启动服务
./tmp_appname -d=true

# 查看进程
ps aux|grep tmp_appname|grep -Ev 'grep\s'|grep --color=auto -E "tmp_appname"
netstat -ant|grep --color=auto tmp_port
echo " 🚗 服务已开启，更多内容请访问 http://localhost:tmp_port"
echo