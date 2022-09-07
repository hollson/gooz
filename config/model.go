package config

import (
	"fmt"
)

// 应用
type app struct {
	Name    string `toml:"name"`    // 应用名称
	Model   string `toml:"model"`   // dev:开发; test:测试; prod:生产
	Lang    string `toml:"lang"`    // 语言环境,zh-CN,en-US等(默认zh-CN)
	Version string `toml:"version"` // 版本号
}

func (p *app) String() string {
	return fmt.Sprintf("%+v", *p)
}

// ===========================================================================

type Authorize struct {
	Account  string `toml:"account"`  // 账号
	Password string `toml:"password"` // 密码
}

// 系统
type system struct {
	Domain   string            `toml:"domain"`   // 官方域名
	Official int               `toml:"official"` // 官方账号
	Monitor  bool              `toml:"monitor"`  // 开启监控
	Robots   []int             `toml:"robots"`   // 机器人账号
	Nodes    map[string]string `toml:"nodes"`    // 服务器节点
	Alert    Authorize         `toml:"alert"`    // 告警账号
}

func (p *system) String() string {
	return fmt.Sprintf("%+v", *p)
}

// ===========================================================================

// 数据库
type service struct {
	Host  string `toml:"host"`  // 主机
	Port  int    `toml:"port"`  // 端口
	Token string `toml:"token"` // 令牌
	Mark  string `toml:"mark"`  // 备注
}

func (p *service) String() string {
	return fmt.Sprintf("%+v", *p)
}

// ===========================================================================

type logFile struct {
	Path     string `toml:"path"`     // 日志输出路径
	Rotation int    `toml:"rotation"` // 文件切割周期(天),默认:1天
	Monitor  int    `toml:"monitor"`  // 文件监控周期(秒),默认:5s
	Retain   int    `toml:"retain"`   // 日志保留时间(天),默认:7天
	Limit    int    `toml:"limit"`    // 单个文件大小上限(M),默认:100M
}
type elastic struct {
	Urls     []string `toml:"urls"`     // 集群
	Account  string   `toml:"account"`  // 账号
	Password string   `toml:"password"` // 密码
	Index    string   `toml:"index"`    // 索引名称
}

// 日志
type log struct {
	Level   string   `toml:"level"`   // error => info => trace
	Console bool     `toml:"console"` // 开启控制台输出(默认开启)
	File    *logFile `toml:"file"`    // 文件日志
	Elastic *elastic `toml:"elastic"` // ELK日志
}

func (p *logFile) String() string {
	return fmt.Sprintf("%+v", *p)
}
func (p *elastic) String() string {
	return fmt.Sprintf("%+v", *p)
}
func (p *log) String() string {
	return fmt.Sprintf("%+v", *p)
}

// ===========================================================================

type database struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Schema   string `toml:"schema"`
	Charset  string `toml:"charset"`
}

// 数据库
type repo struct {
	Master database   `toml:"master"` // 主库
	Slave  []database `toml:"slave"`  // 从库
}

func (p *database) String() string {
	return fmt.Sprintf("%+v", *p)
}
func (p *repo) String() string {
	return fmt.Sprintf("%+v", *p)
}

// ===========================================================================

// Redis数据库
type redis struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
	Auth string `toml:"auth"`
}

func (p *redis) String() string {
	return fmt.Sprintf("%+v", *p)
}

// ===========================================================================

// kafka
type kafka struct {
	Hosts    []string `toml:"hosts"`
	Account  string   `toml:"account"`
	Password string   `toml:"password"`
	Producer string   `toml:"producer_topic"`
	Consumer string   `toml:"consumer_topic"`
	Group    string   `toml:"group"`
	Sasl     bool     `toml:"sasl_enable"`
}

func (p *kafka) String() string {
	return fmt.Sprintf("%+v", *p)
}
