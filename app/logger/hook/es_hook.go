// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

/**************************************
  ElasticSearch日志输出
**************************************/
package hook

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var _client *elastic.Client

type EsHook struct {
	Urls      []string     // es集群
	Account   string       // 用户名
	Password  string       // 密码
	IndexName string       // 索引名称
	Level     logrus.Level // 日志级别
}

// 创建异步ES-Hook对象
func (e *EsHook) NewHook() (hook logrus.Hook, err error) {
	var elasticSetBasicAuth elastic.ClientOptionFunc = nil
	if strings.TrimSpace(e.Account) != "" {
		elasticSetBasicAuth = elastic.SetBasicAuth(e.Account, e.Password)
	}

	_client, err = elastic.NewClient(
		elastic.SetURL(e.Urls...),
		elasticSetBasicAuth,
		elastic.SetSniff(false),                        // 启用或禁用嗅探器(默认启用)
		elastic.SetHealthcheckInterval(30*time.Second), // 设置两次运行状况检查之间的间隔。默认60秒
		// elastic.SetErrorLog(log.New(os.Stderr, "ERR:", log.LstdFlags)),
		// elastic.SetInfoLog(log.New(os.Stdout, "INF:", log.LstdFlags)),
		// elastic.SetTraceLog(log.New(os.Stdout, "TRC:", log.LstdFlags)),
	)
	if err != nil {
		return nil, fmt.Errorf("create elastic client error: %v ,eshook: %+v", err, e)
	}
	hook, err = NewAsyncElasticHook(_client, "localhost", e.Level, e.IndexName)
	return
}

// 检查Elastic连通性
func (e *EsHook) PingElastic() (info *elastic.PingResult, httpStatus int, err error) {
	if _client == nil {
		return nil, 0, errors.New("elastic client is null")
	}
	info, httpStatus, err = _client.Ping(e.Urls[0]).Do(context.Background())
	if err != nil {
		logrus.Printf("ping es failed, err: %v", err)
		return nil, 0, err
	}
	return
}
