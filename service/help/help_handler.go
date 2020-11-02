// ------------------------------------------------------------------
// @ Author: hollson <hollson@live.cn>
// @ Date: 2019-12-01
// @ Version: 1.0.0
// 帮助中心：提供调试、ping和帮助查询
// -------------------------------------------------------------------------------------

package help

import (
    "expvar"
    "fmt"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/hollson/gooz/app/config"
)

// GetCurrentRunningStats 返回当前运行信息
func RuntimeStatHandler(c *gin.Context) {
    c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
    first := true
    report := func(key string, value interface{}) {
        if !first {
            fmt.Fprintf(c.Writer, ",\n")
        }
        first = false
        if str, ok := value.(string); ok {
            fmt.Fprintf(c.Writer, "%q: %q", key, str)
        } else {
            fmt.Fprintf(c.Writer, "%q: %v", key, value)
        }
    }

    fmt.Fprintf(c.Writer, "{\n")
    expvar.Do(func(kv expvar.KeyValue) {
        report(kv.Key, kv.Value)
    })
    fmt.Fprintf(c.Writer, "\n}\n")
    c.String(http.StatusOK, "")
}

// 复查配置文件
func ConfigViewHandler(c *gin.Context) {
    if os.Getenv("GOOZ_Env") == "prod" {
        c.String(200, "PROD")
        return
    }
    c.JSON(200, gin.H{
        "App":   *config.App,
        "PG":    *config.Postgres,
        "Redis": *config.Log,
        "Mysql": *config.Mysql,
    })
}

// 复查配置文件,fixme:数据库，redis...
func PingHandler(c *gin.Context) {

    c.JSON(200, gin.H{
        "MySQL":  "Pong",
        "PgSQL":  "Pong",
        "Redis":  "Pong",
        "Etcd":   "Pong",
        "Zipkin": "Pong",
    })
}
