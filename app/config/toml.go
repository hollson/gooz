package config

// 用于初始化系统配置文件，须优先维护。
// 配置项优先级： 项目配置 < 环境变量 < ETC配置 < Auto配置
var TPL = `# 系统配置文件（可根据规则删除不需要的配置项）
# %s

# =============================Application==================================
# 应用程序配置
# 说明：物理主机的系统环境变量优先级将高于当前的配置（如：GOENVIRONMENT）。
# ==========================================================================
[app]
	appName = "Gooz"
	port = ":8888"
	environment = "dev"   # dev:开发、test:测试、prod:生产


# ================================Logger====================================
# 日志配置
# ==========================================================================
[log]
	level = "info"      # error => info => trance
	console = true      # 开启控制台输出(默认开启)
[log.file]
	path="./logs"       # 日志输出路径
	rotation=1          # 文件切割周期(天),默认：1天
	monitor=5   		# 文件监控周期(秒)，默认：5s
	retain=90		    # 日志保留时间(天)，默认：7天
	limit_size=200		# 单个文件大小上限(M)，默认：100M
[log.elastic]
	urls= ["http://10.0.0.13:9200","http://127.0.0.1:920"]
    account="elastic"   # 账户名
    password="123456"   # 账号
    index="gooz-server" # 密码


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
