package stats

import (
	"expvar"
	"sync"

	"github.com/gin-gonic/gin"
)

var apiVisit ApiVisit

type ApiVisit struct {
	visit map[string]int
	mu    sync.RWMutex
}

func apiStats() interface{} {
	apiVisit.mu.RLock()
	defer apiVisit.mu.RUnlock()
	return apiVisit.visit
}

func ApiVisitHandler(c *gin.Context) {
	if _, ok := apiVisit.visit[c.Request.RequestURI]; ok {
		apiVisit.mu.Lock()
		defer apiVisit.mu.Unlock()
		apiVisit.visit[c.Request.RequestURI] += 1
		return
	}
	apiVisit.visit[c.Request.RequestURI] = 1
}

func init() {
	apiVisit = ApiVisit{visit: make(map[string]int)}
	expvar.Publish("API统计", expvar.Func(apiStats))
}
