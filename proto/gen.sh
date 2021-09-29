#!/usr/bin/env bash

rm -rf ./_gen
mkdir -p ./_gen/go ./_gen/py ./_gen/java ./_gen/cs

#proto2语法规范参考：
#https://developers.google.com/protocol-buffers/docs/reference/proto2-spec

#proto3语法规范参考：
#https://developers.google.com/protocol-buffers/docs/reference/proto3-spec

# 编译proto文件：
# include包含了官方模版
# 格式：protoc [OPTION] PROTO_FILES
# 参数：-I 指定一组ptoto文件

##注意事项(libprotoc 3.9.0)：
# 1. C#只支持proto3语法
# 2. PHP只支持proto3语法

#protoc -I ./ \
#--go_out=./_gen/go \
#--java_out=./_gen/java/ \
#--python_out=./_gen/py/ \
#./*.proto


# 多语言版本
protoc -I ./ \
--go_out=plugins=grpc:../gen/go \
--csharp_out=./_gen/cs/ \
--java_out=./_gen/java/ \
./*.proto
