#!/bin/bash

: '说明：在PPGo_job中配置定时任务，执行以下示例的curl请求参数，
    即：在url末尾追加电子签名，在服务端进行签名验证，以防止外部恶意调用。
    请求方法：md5(secret.timespan)'

kill -9 $(pidof deeplink-linux-amd64-v2.0.1)
./deeplink-linux-amd64-v2.0.1 -d=true
ps -ef|grep deeplink
