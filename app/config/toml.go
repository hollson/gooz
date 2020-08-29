package config

// 用于初始化系统配置文件，须优先维护。
var TPL = `# 系统配置文件（可根据规则删除不需要的配置项）
# %s

# =============================Application==================================
# 应用程序配置
# 说明：物理主机的系统环境变量优先级将高于当前的配置（如：GOENVIRONMENT）。
# ==========================================================================
[app]
	appName = "Deeplink"
	port = ":8080"
	version = "v1.0.1"
	environment = "dev"   # dev:开发、test:测试、prod:生产


# ================================Logger====================================
# 日志配置
# ==========================================================================
[log]
	path = "./logs"
	level = "error"
	hook = ""


# =================================Mysql====================================
# Mysql数据库
# ==========================================================================
[mysql]
	host = "127.0.0.1"
	port = 3306
	user = "root"
	password = "123456"
	schema = "gooz"
	charset = "utf8"


# ==============================Postgresql==================================
# PG数据库集群
# ==========================================================================
[[postgres]]
	host = "127.0.0.1"
	port = 5432
	user = "postgres"
	password = "123456"
	schema = "gooz"
	sslmode = "disable"
[[postgres]]
	host = "127.0.0.1"
	port = 5433
	user = "postgres"
	password = "123456"
	schema = "gooz"
	sslmode = "disable"
[[postgres]]
	host = "127.0.0.1"
	port = 5434
	user = "postgres"
	password = "123456"
	schema = "gooz"
	sslmode = "disable"


# =================================Redis====================================
# Redis哨兵/集群；以下单例、哨兵、集群配置没有关联性
# ==========================================================================
[redis]
	host = "127.0.0.1"
	port = 26379
	#auth = "123456"

[sentinel]
	[[sentinel.member]]
		host = "127.0.0.1"
		port = 26379
		#auth = "123456"
	[[sentinel.member]]
		host = "127.0.0.1"
		port = 26380
		#auth = "123456"
	[[sentinel.member]]
		host = "127.0.0.1"
		port = 26381
		#auth = "123456"

[cluster]
	[[cluster.member]]
		host = "127.0.0.1"
		port = 1001
	[[cluster.member]]
		host = "127.0.0.1"
		port = 1002
`
