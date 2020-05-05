#当前操作系统
GOOS=$(shell go env GOOS)

AppName="Deeplink" #应用名称
VERSION="v1.0.1"   #版本号
CGO=0			   #是否开启Cgo，0：不开启，1：开启


## all@编译并运行(默认项，可直接执行make命令)。
all: build run


## build@根据系统类型交叉编译(支持linux、darwin和windows)。
.PHONY:build
build: clean
	@echo "\033[34m开始编译...\033[0m"
	@if [ $(GOOS) = "linux" ]; \
	then \
		echo "\033[35m当前系统类型：linux\033[0m"; \
		CGO_ENABLED=$(CGO) GOOS=linux GOARCH=amd64 go build -x -o ./bin/"`echo $(AppName)|sed s/[[:space:]]//g`-linux-amd64-$(VERSION)"; \
	elif [ $(GOOS) = "darwin" ]; \
	then \
		echo "\033[35m当前系统类型：darwin\033[0m"; \
		CGO_ENABLED=$(CGO) GOOS=darwin GOARCH=amd64 go build -x -o ./bin/"`echo $(AppName)|sed s/[[:space:]]//g`-darwin-amd64-$(VERSION)"; \
	elif [ $(GOOS) = "windows" ]; \
	then \
		echo "\033[35m当前系统类型：windows\033[0m"; \
		CGO_ENABLED=$(CGO) GOOS=windows GOARCH=amd64 go build -x -o ./bin/"`echo $(AppName)|sed s/[[:space:]]//g`-win-amd64-$(VERSION).exe"; \
	else \
		echo "未知的操作系统类型."; \
	fi
	echo "\033[35m编译完成\033[0m"; \


## clean@清理编译、日志和缓存等数据。
.PHONY:clean
clean:
	@echo "\033[31m开始清理...\033[0m";
	@rm -rf ./bin;
	@rm -rf ./logs;
	@rm -rf ./log;
	@rm -rf ./cache;
	@rm -rf ./pid;


## deploy@发布到远程Web服务器。
.PHONY:deploy
deploy:
	@#压缩本地发布包,并推送到远程服务器
	@echo "\033[0;32m发布中...\033[0m"
	tar -zcvf $(AppName)-release-$(VERSION)-tar.gz public
	scp $(AppName)-release-$(VERSION)-tar.gz root@www.xxx.com:/srv/www/$(AppNAme)

	@#执行远程命令,进行解压、重启，并清理本地压缩包
	echo -e "\033[0;32m执行远程清理...\033[0m"
	ssh root@www.mafool.com 'rm -rf /srv/www/$(AppName)'
	ssh root@www.mafool.com 'cd /srv/www/$(AppName) && tar -zxvf $(AppName)-release-$(VERSION)-tar.gz && nginx -s reload'
	rm -f mafool-blog.tar.gz


## commit <msg>@Git本地Commit(如:make commit msg="备注内容",msg参数为可选项)。
.PHONY:commit
message:=$(if $(msg),$(msg),"Rebuilded at $$(date '+%Y年%m月%d日 %H时%M分%S秒')");
commit:
	@echo "\033[0;34mPush to remote...\033[0m"
	@git add .
	@git commit -m $(message)
	@echo "\033[0;31mCommit成功\033[0m"


## push <msg>@执行commit并push到远程Git仓库,格式如commit命令。
.PHONY:push
push:commit
	@git push #origin master
	@echo "\033[0;31mPush成功\033[0m"


## proto@更新并编译proto文件。
.PHONY:proto
proto:
	@echo "更新并编译proto文件";
	#git pull
	#protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/user/user.proto


## run@运行(可从命令行接收参数,如:make run daemon=true)。
.PHONY:run
run:
	@go run main.go $(deamon)

## update@更新Git和Submodule。
.PHONY:update
update:
	@git submodule update --init --recursive;


## xorm@更新数据库模型。
.PHONY:xorm
xorm:
	@echo "更新数据库模型";


## help@查看make帮助。
.PHONY:help
help:Makefile
	@echo "Usage:\n  make [command]"
	@echo
	@echo "Available Commands:"
	@sed -n "s/^##//p" $< | column -t -s '@' |grep --color=auto "^[[:space:]][a-z]\+[[:space:]]"
	@echo
	@echo "For more to see https://makefiletutorial.com/"
