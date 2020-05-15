# App基本信息
AppName="Deeplink"	#应用名称
VERSION="v1.0.1"	#版本号
CGO=0				#是否开启Cgo，0：不开启，1：开启


## all@可选的命令参数，执行build和run命令。
all: proto build


## build <os>@编译(格式：make build os=linux/darwin/windows,os为可选参数)。
.PHONY:build
OS:=$(if $(os),$(os),$(shell go env GOOS))
ARCH:=$(if $(arch),$(arch),"amd64")
build: clean
	@echo "\033[34m 😊 开始编译...\033[0m"
	@if [ $(OS) = "linux" ]; \
	then \
		echo "\033[35m 🍵 编译环境：$(OS)\033[0m"; \
		AppName="`echo $(AppName)-$(OS)-$(ARCH)-$(VERSION)|sed s/[[:space:]]//g`";\
		CGO_ENABLED=$(CGO) GOOS=$(OS) GOARCH=$(ARCH) go build -o ./tmp/$${AppName}; \
		cp -rp ./conf ./tmp && cp ./scripts/run.sh ./tmp && cp ./scripts/stop.sh ./tmp; \
		sed -i "s/tmp_appname/$${AppName}/g" ./tmp/run.sh; \
		sed -i "s/tmp_appname/$${AppName}/g" ./tmp/stop.sh; \
		echo "\033[35m ✅  编译完成\033[0m";\
		echo "输出路径：./tmp" && ls -hl ./tmp;\
	elif [ $(OS) = "darwin" ]; \
	then \
		echo "\033[35m 🍵 编译环境：$(OS)\033[0m"; \
		AppName="`echo $(AppName)-$(OS)-$(ARCH)-$(VERSION)|sed s/[[:space:]]//g`";\
		CGO_ENABLED=$(CGO) GOOS=$(OS) GOARCH=$(ARCH) go build -o ./tmp/$${AppName}; \
		cp -rp ./conf ./tmp && cp ./scripts/run.sh ./tmp && cp ./scripts/stop.sh ./tmp; \
		sed -i "" "s/tmp_appname/$${AppName}/g" `grep -rl tmp_appname ./tmp/run.sh`; \
		sed -i "" "s/tmp_appname/$${AppName}/g" `grep -rl tmp_appname ./tmp/stop.sh`; \
		echo "\033[35m ✅  编译完成\033[0m";\
		echo "输出路径：./tmp" && ls -hl ./tmp;\
	elif [ $(OS) = "windows" ]; \
	then \
		echo "\033[35m 🍵 编译环境：windows\033[0m"; \
		CGO_ENABLED=$(CGO) GOOS=windows GOARCH=amd64 go build -o ./tmp/"`echo $(AppName)-win-amd64-$(VERSION).exe|sed s/[[:space:]]//g`"; \
		cp -rp ./conf ./tmp \
        echo "\033[35m ✅  编译完成\033[0m";\
        echo "输出路径：./tmp" && ls -hl ./tmp;\
	else \
		echo " ❌  未知的操作系统类型:$(OS)."; \
	fi;


## clean@清理编译、日志和缓存等数据。
.PHONY:clean
clean:
	@rm -rf ./bin;
	@rm -rf ./logs;
	@rm -rf ./log;
	@rm -rf ./cache;
	@rm -rf ./pid;
	@rm -rf ./release;
	@rm -rf ./debug;
	@rm -rf ./tmp;
	@rm -rf ./temp;
	@echo "\033[31m ✅  清理完成\033[0m";


## commit <msg>@提交Git(格式:make commit msg="备注内容",msg为可选参数)。
.PHONY:commit
message:=$(if $(msg),$(msg),"Rebuilded at $$(date '+%Y年%m月%d日 %H时%M分%S秒')")
commit:
	@echo "\033[0;34mPush to remote...\033[0m"
	@git add .
	@git commit -m $(message)
	@echo "\033[0;31m 💿 Commit成功\033[0m"


## deploy@[远程]发布到远程服务器。
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
	@echo "\033[31m ✅  发布完成\033[0m";


## install@[本地]安装并启动服务。
.PHONY:install
install:
	@pkill $(AppName)
	@sudo cp -rp ./release /tmp/ && mv /tmp/;
	@/tmp/xxxx -d=true
	@echo "\033[31m ✅  服务已启动\033[0m";
	@ps aux|grep $(AppName)


## push <msg>@提交并推送到Git仓库(格式:make push msg="备注内容",msg为可选参数)。
.PHONY:push
push:commit
	@git push #origin master
	@echo "\033[0;31m ⬆️ Push成功\033[0m"


## proto@更新并编译proto文件。
.PHONY:proto
proto:
	@cd proto && ./gen.sh && cd -;
	@echo "\033[35m ✅  Proto编译完成\033[0m"; \


## run@运行服务。
.PHONY:proto run
run: clean
	@echo " ⚽  启动服务..."
	@go run main.go $(deamon)


## update@更新Git和Submodule。
.PHONY:update
update:
	@git submodule update --init --recursive;


## xorm@根据数据表结构生成实体,支持mysql、postgres、sqlite等。
.PHONY:xorm
Templates=$(GOPATH)/src/github.com/go-xorm/cmd/xorm/templates/goxorm/
REPO_PATH=$$(pwd)/repo
xorm:
	@sudo rm -rf $(REPO_PATH)/models/*;
	@#sudo xorm reverse mysql root:"123456"@"(127.0.1:3306)"/demo?charset=utf8 $(Templates) $(REPO_PATH)/models;
	@sudo xorm reverse postgres "user=postgres password=123456 dbname=testdb host=127.0.0.1 port=5432 sslmode=disable" $(Templates) $(REPO_PATH)/models;
	@#xorm reverse sqite3 test.db templates/goxorm C:\temp
	@echo "\033[31m ✅  Reverse完成\033[0m";
#https://pkg.go.dev/github.com/lib/pq?tab=doc


## help@查看make帮助。
.PHONY:help
help:Makefile
	@echo "Usage:\n  make [command]"
	@echo
	@echo "Available Commands:"
	@sed -n "s/^##//p" $< | column -t -s '@' |grep --color=auto "^[[:space:]][a-z]\+[[:space:]]"
	@echo
	@echo "For more to see https://makefiletutorial.com/"