#当前操作系统
GOOS=$(shell go env GOOS)

AppName="Deeplink" #应用名称
VERSION="v1.0.1"   #版本号
CGO=0			   #是否开启Cgo，0：不开启，1：开启

#CMD: make		-- 编译并运行
all: build run


#CMD: make build	-- 编译
.PHONY:build
build: clean
	@echo "开始编译..."
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


#CMD: make run		-- 运行(可从命令行接收参数，如：make run daemon=true)
.PHONY:run
run:
	@go run main.go $(deamon)

#CMD: make push <msg>	-- 推送到远程仓库(msg为空时，使用时间标记的默认注释)
.PHONY:push
push:
	@echo -e "\033[0;32mPush to GitHub...\033[0m"
	git add .
	msg="rebuilding site $(date)"
	@if [ $# -eq 1 ]; then \
	  msg="$1" ; \
	fi
	git commit -m "$msg"
	git push origin master
	echo "源码推送成功！"

#CMD: make deploy	-- 发布
.PHONY:deploy
deploy:
	@echo "\033[0;32m发布中...\033[0m"

	#编译
#	rm -rf ./public/*
#	hugo -v --baseURL=http://www.mafool.com
#	echo -e "\033[0;32m编译完成...\033[0m"
#
#	# 打包上传
#	rm -f mafool-blog.tar.gz
#	tar -zcvf mafool-blog.tar.gz public
#	scp mafool-blog.tar.gz root@www.mafool.com:/srv/www/
#
#	echo -e "\033[0;32m执行远程清理...\033[0m"
#	ssh root@www.mafool.com 'rm -rf /srv/www/mafool-blog'
#	ssh root@www.mafool.com 'cd /srv/www && tar -zxvf mafool-blog.tar.gz && mv public mafool-blog && nginx -s reload'
#	rm -f mafool-blog.tar.gz


#CMD: make proto	-- 更新并编译proto文件
.PHONY:proto
proto:
	@echo "更新并编译proto文件";
	#git pull
	#protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/user/user.proto


#CMD: make xorm		-- 更新数据库模型
.PHONY:xorm
xorm:
	@echo "更新数据库模型";


#CMD: make clean	-- 清理编译、日志和缓存等数据
.PHONY:clean
clean:
	@echo "开始清理..."
	@rm -rf ./bin;
	@rm -rf ./logs;
	@rm -rf ./log;
	@rm -rf ./cache;
	@rm -rf ./pid;


#CMD: make help		-- 查看make帮助
.PHONY:help
help:
	grep -ie "^#CMD" ./Makefile