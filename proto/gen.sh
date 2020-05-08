#!/usr/bin/env bash

rm -rf ./_gen
mkdir -p ./_gen/go ./_gen/py ./_gen/php ./_gen/cs ./_gen/java

#proto2语法规范参考：
#https://developers.google.com/protocol-buffers/docs/reference/proto2-spec
#proto3语法规范参考：
#https://developers.google.com/protocol-buffers/docs/reference/proto3-spec

##注意事项(libprotoc 3.9.0)：
# 1. C#只支持proto3语法
# 2. PHP只支持proto3语法


:'
功能：执行编译命令
用法：protoc [OPTION] PROTO_FILES
参数：-I 指定一组ptoto文件
'
protoc -I ./ \
--go_out=plugins=grpc:./_gen/go \
--java_out=./_gen/java/ \
--python_out=./_gen/py/ \
./*.proto

# 多语言版本
#protoc -I ./ \
#--go_out=plugins=grpc:./_gen/go \
#--csharp_out=./_gen/cs/ \
#--java_out=./_gen/java/ \
#--php_out=./_gen/php/ \
#--python_out=./_gen/py/ \
#./*.proto
