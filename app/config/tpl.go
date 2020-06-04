package config

var TPL = `# Generated by the initial configuration template .
# %s

# 应用程序信息
[app]
    appName = "Deeplink"  #应用名称
    port = ":8080"        #服务端口
    version = "v1.0.1"    #版本号
    environment = "dev"   #运行环境(dev:开发、test:测试、prod:生产)


[log]
    path = "./logs"
    level = "error"
    hook = ""


# Mysql数据库
[mysql]
    enable = true
    host = "127.0.0.1"
    port = 3306
    user = "root"
    password = "123456"
    schema = "deeplink"
    charset = "utf8"


# PG数据库集群
[[postgres]]
    enable = true
    host = "127.0.0.1"
    port = 5432
    user = "postgres"
    password = "123456"
    schema = "testdb"
    sslmode = "disable"
[[postgres]]
    enable = true
    host = "127.0.0.1"
    port = 5433
    user = "postgres"
    password = "123456"
    schema = "testdb"
    sslmode = "disable"
[[postgres]]
    enable = true
    host = "127.0.0.1"
    port = 5434
    user = "postgres"
    password = "123456"
    schema = "testdb"
    sslmode = "disable"


#Redis哨兵/集群
[redis]
    [redis.master]
        host = "127.0.0.1"
        port = 6379
    [redis.cluster1]
        host = "127.0.0.1"
        port = 6380
`
