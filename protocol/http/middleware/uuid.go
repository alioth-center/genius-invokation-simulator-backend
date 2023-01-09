package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sunist-c/genius-invokation-simulator-backend/util"
)

// GetUUID 获取context中携带的uuid信息，若没有则生成并写入一个uuid
func GetUUID(ctx *gin.Context, uuidKey string) (uuid string) {
	if result, gotten := ctx.Get(uuidKey); !gotten {
		uuid = util.GenerateUUID()
		ctx.Set(uuidKey, uuid)
		return uuid
	} else {
		return result.(string)
	}
}

// NewUUIDTagger 创建一个UUID标记器，将会往context中写入uuid
func NewUUIDTagger(uuidKey string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Set(uuidKey, util.GenerateUUID())
	}
}
