# This file is a template, and might need editing before it works on your project.
# https://blog.csdn.net/Gusand/article/details/100658595

FROM alpine
MAINTAINER holson hollson@live.com
WORKDIR /app
COPY . .
EXPOSE 8080

#RUN find . -name *.go -exec rm -rf {} \ #删除源文件
#RUN find -type d -empty |xargs rm -rf;  #删除空目录
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

CMD ./gooz
