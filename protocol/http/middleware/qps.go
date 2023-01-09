package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/kv"
	"time"
)

// NewQPSLimiter 生成一个QPS限制器，以IP来源为基础
func NewQPSLimiter(limit time.Duration, ipKey string) func(ctx *gin.Context) {
	var limiter = kv.NewSyncMap[time.Time]()
	return func(ctx *gin.Context) {
		if success, ip := GetIPTrace(ctx, ipKey); !success {
			// 无法成功获取客户端IP，返回BadRequest
			ctx.AbortWithStatus(400)
		} else {
			if limiter.Exists(ip) {
				if t := limiter.Get(ip); !t.Add(limit).Before(time.Now()) {
					// 请求过快，返回PreconditionFailed
					ctx.AbortWithStatus(412)
				} else {
					// 正常请求，更新访问时间
					limiter.Set(ip, time.Now())
				}
			} else {
				// 之前未访问过，记录访问时间
				limiter.Set(ip, time.Now())
			}
		}
	}
}
