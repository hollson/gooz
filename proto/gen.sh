#!/usr/bin/env bash

rm -rf ./_gen
mkdir -p ./_gen/go ./_gen/py ./_gen/php ./_gen/cs ./_gen/java

: '
功能：生成相应编程语言的protobuf文件
用法：protoc [OPTION] PROTO_FILES
参数：-I 指定一组ptoto文件
'
protoc -I ./ \
--go_out=plugins=grpc:./_gen/go \
--csharp_out=./_gen/cs/ \
--java_out=./_gen/java/ \
--php_out=./_gen/php/ \
--python_out=./_gen/py/ \
./*.proto

echo " ✅ proto文件编译完成！！！"