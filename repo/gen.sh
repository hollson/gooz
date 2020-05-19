#!/usr/bin/env bash

# XORM参考文档：
# https://xorm.io/
# https://gitea.com/xorm/xorm/src/branch/master/README_CN.md

# 从数据库生成对象代码
cd .. && make xorm && cd -


# https://pkg.go.dev/github.com/lib/pq?tab=doc#StringArray