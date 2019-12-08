#!/bin/bash

: '说明：在PPGo_job中配置定时任务，执行以下示例的curl请求参数，
    即：在url末尾追加电子签名，在服务端进行签名验证，以防止外部恶意调用。
    请求方法：md5(secret.timespan)'

secret="204NAOB7JND0YRRA"       #密钥
timespan=$(date +%s)            #时间戳
raw=$secret"."$timespan         #签名字符串
sign=`echo $raw| md5sum | cut -d ' ' -f 1`  #签名
attach="timespan="$timespan"&sign="$sign

# api接口（注意末尾是否要加&）
api_url="http://deeplink.xxx.com/api/update?over=0&"$attach
echo $api_url
curl $api_url