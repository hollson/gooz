# Gooz
基于Gin+Zk+Xorm+Mysql+Redis+Zipkin的Restfull+Jwt的项目模板


## 配置管理
> 基于TOML v0.4.0的配置管理
> 
> Toml语法可查看[Toml官方文档](https://github.com/toml-lang/toml)或[中文文档](
> https://github.com/toml-lang/toml/blob/master/versions/cn/toml-v0.4.0.md)。
>
> GO客户端：[Toml-Go客户端](https://github.com/BurntSushi/toml)，[使用示例](https://github.com/BurntSushi/toml/tree/master/_examples)。

```bash
# 安装toml-go客户端
go get github.com/BurntSushi/toml

# 验证toml语法
go get github.com/BurntSushi/toml/cmd/tomlv
tomlv some-toml-file.toml
```

# 目录结构
```txt
├── apiauth
├── app
│   ├── app.go
│   ├── auth
│   ├── cmd.go
│   ├── config
│   ├── etcd
│   ├── export
│   ├── logger
│   ├── midware
│   └── router.go
├── asset
│   └── index.html
├── base
│   └── api.go
├── cmd
├── conf
├── docs
├── main.go
├── proto
├── repo
├── service
│   ├── account
│   ├── article
│   ├── help
│   └── home
├── setup
├── util
└── version.go
```


# 配置文件

## Zipkin链路追踪
https://segmentfault.com/a/1190000016677230




xorm问题
https://www.cnblogs.com/nickchou/p/9561561.html

https://www.jianshu.com/p/b15cdc114458


## 代码规范
https://github.com/uber-go/guide
https://yq.aliyun.com/articles/598076?spm=a2c4e.11153940.0.0.2ef65e58ohXosw
https://yq.aliyun.com/articles/603241?utm_content=m_1000003274

# 优雅启动
https://github.com/fvbock/endless


// FIXME + 说明：
// +build sss
