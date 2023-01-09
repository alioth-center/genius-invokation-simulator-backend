package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/util"
	"strconv"
	"strings"
)

func GetIPTrace(ctx *gin.Context, traceKey string) (has bool, ip uint) {
	if result, gotten := ctx.Get(traceKey); !gotten {
		return false, 0
	} else if ipResult, ok := result.(uint); !ok {
		return false, 0
	} else {
		return has, ipResult
	}
}

func ConvertIPToUint(ip string) (success bool, result uint) {
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

func ConvertUintToIP(ip uint) (result string) {
	bytes := make([]byte, 4)
	for i := 0; i < 4; i++ {
		bytes[i] = byte(ip % 256)
		ip = ip >> 8
	}
	return fmt.Sprintf("%d.%d.%d.%d", bytes[0], bytes[1], bytes[2], bytes[3])
}

func NewIPTracer(traceKey string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if success, ip := ConvertIPToUint(util.GetClientIP(ctx)); !success {
			// 无法成功获取客户端IP，返回BadRequest
			ctx.AbortWithStatus(400)
		} else {
			ctx.Set(traceKey, ip)
		}
	}
}
