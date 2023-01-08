package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/kv"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/util"
	"strconv"
	"strings"
	"time"
)

func convertIPToUint(ip string) (success bool, result uint) {
	splits := strings.Split(ip, ".")
	if len(splits) != 4 {
		return false, 0
	} else {
		for _, s := range splits {
			if intResult, err := strconv.Atoi(s); err != nil {
				return false, 0
			} else if intResult > 256 || intResult < 0 {
				return false, 0
			} else {
				result = result<<8 + uint(intResult)
			}
		}
		return true, result
	}
}

func convertUintToIP(ip uint) (result string) {
	bytes := make([]byte, 4)
	for i := 0; i < 4; i++ {
		bytes[i] = byte(ip % 256)
		ip = ip >> 8
	}
	return fmt.Sprintf("%d.%d.%d.%d", bytes[0], bytes[1], bytes[2], bytes[3])
}

// NewQPSLimiter 生成一个QPS限制器，以IP来源为基础
func NewQPSLimiter(limit time.Duration) func(ctx *gin.Context) {
	var limiter = kv.NewSyncMap[time.Time]()
	return func(ctx *gin.Context) {
		if success, ip := convertIPToUint(util.GetClientIP(ctx)); !success {
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
