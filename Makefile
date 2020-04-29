#当前操作系统
GOOS=$(shell go env GOOS)

AppName="Deeplink" #应用名称
VERSION="v1.0.1"   #版本号
CGO=0			   #是否开启Cgo，0：不开启，1：开启

# 编译并运行
all: build run


# 编译
.PHONY:build
build:
	@if [ $(GOOS) = "linux" ]; \
	then \
		echo "当前系统类型：linux"; \
		CGO_ENABLED=$(CGO) GOOS=linux GOARCH=amd64 go build -x -o ./bin/"`echo $(AppName)|sed s/[[:space:]]//g`-linux-amd64-$(VERSION)"; \
	elif [ $(GOOS) = "darwin" ]; \
	then \
		echo "当前系统类型：darwin"; \
		CGO_ENABLED=$(CGO) GOOS=darwin GOARCH=amd64 go build -x -o ./bin/"`echo $(AppName)|sed s/[[:space:]]//g`-darwin-amd64-$(VERSION)"; \
	elif [ $(GOOS) = "windows" ]; \
	then \
		echo "当前系统类型：windows"; \
		CGO_ENABLED=$(CGO) GOOS=windows GOARCH=amd64 go build -x -o ./bin/"`echo $(AppName)|sed s/[[:space:]]//g`-win-amd64-$(VERSION).exe"; \
	else \
		echo "未知的操作系统类型."; \
	fi


# 运行(可从命令行接收参数，如：make run os=linux)
.PYONY:run
run:
	@echo "运行";

# 发布
.PYONY:deploy
deploy:
	@echo "发布";

# 更新并编译proto文件
.PHONY:proto
proto:
	@echo "更新并编译proto文件";
	#git pull
	#protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/user/user.proto


# 更新数据库模型
.PHONY:xorm
xorm:
	@echo "更新数据库模型";


# 清理编译、日志和缓存等数据
.PHONY:clean
clean:
	rm -rf ./bin;
	rm -rf ./logs;
	rm -rf ./log;
	rm -rf ./cache;