package util

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func convertToInt(source string) (success bool, result int) {
	if intResult, err := strconv.Atoi(source); err != nil {
		return false, 0
	} else {
		return true, intResult
	}
}

func convertToFloat(source string) (success bool, result float64) {
	if floatResult, err := strconv.ParseFloat(source, 64); err != nil {
		return false, float64(0)
	} else {
		return true, floatResult
	}
}

func convertToBool(source string) (success bool, result bool) {
	if boolResult, err := strconv.ParseBool(source); err != nil {
		return false, false
	} else {
		return true, boolResult
	}
}

func convertToStruct[T any](source string) (success bool, result T) {
	if err := json.Unmarshal([]byte(source), &result); err != nil {
		return false, result
	} else {
		return true, result
	}
}

// BindJson 将请求体中的数据绑定到给定的entity中，entity需要可寻址，失败直接返回BadRequest
func BindJson[T any](ctx *gin.Context, entity T) (success bool) {
	if err := ctx.ShouldBindJSON(entity); err != nil {
		return false
	} else {
		return true
	}
}

// QueryPath 从URL进行查询，支持路径参数与查询参数
func QueryPath(ctx *gin.Context, key string) (has bool, result string) {
	if strings.HasPrefix(key, ":") {
		result, has = ctx.Params.Get(strings.Trim(key, ":"))
	} else {
		result, has = ctx.GetQuery(key)
	}

	return has, result
}

// QueryPathInt 从URL中查询一个int类型的结果，调用QueryPath并进行转化
func QueryPathInt(ctx *gin.Context, key string) (has bool, result int) {
	if gotten, stringResult := QueryPath(ctx, key); !gotten {
		return false, 0
	} else {
		return convertToInt(stringResult)
	}
}

// QueryPathFloat 从URL中查询一个float64类型的结果，调QueryPathQueryPath并进行转化
func QueryPathFloat(ctx *gin.Context, key string) (has bool, result float64) {
	if gotten, stringResult := QueryPath(ctx, key); !gotten {
		return false, float64(0)
	} else {
		return convertToFloat(stringResult)
	}
}

// QueryPathBool 从URL中查询一个bool类型的结果，调QueryPathQueryPath并进行转化
func QueryPathBool(ctx *gin.Context, key string) (has bool, result bool) {
	if gotten, stringResult := QueryPath(ctx, key); !gotten {
		return false, false
	} else {
		return convertToBool(stringResult)
	}
}

// QueryPathJson 从URL中查询一个json对象，调用QueryPath并进行转化
func QueryPathJson[T any](ctx *gin.Context, key string) (has bool, result T) {
	if gotten, stringResult := QueryPath(ctx, key); !gotten {
		return false, result
	} else {
		return convertToStruct[T](stringResult)
	}
}

// QueryCookie 从cookie中获取指定key的值
func QueryCookie(ctx *gin.Context, key string) (has bool, cookie string) {
	if cookieResult, err := ctx.Cookie(key); err != nil {
		return false, cookie
	} else {
		return true, cookieResult
	}
}

// QueryHeaders 从请求头中获取指定key的值
func QueryHeaders(ctx *gin.Context, key string) (has bool, result string) {
	if result = ctx.GetHeader(key); result == "" {
		return false, ""
	} else {
		return true, result
	}
}

// QueryHeadersInt 从请求头中获取一个int类型的值
func QueryHeadersInt(ctx *gin.Context, key string) (has bool, result int) {
	if gotten, stringResult := QueryHeaders(ctx, key); !gotten {
		return false, 0
	} else {
		return convertToInt(stringResult)
	}
}

// QueryHeadersFloat 从请求头中获取一个float类型的值
func QueryHeadersFloat(ctx *gin.Context, key string) (has bool, result float64) {
	if gotten, stringResult := QueryHeaders(ctx, key); !gotten {
		return false, float64(0)
	} else {
		return convertToFloat(stringResult)
	}
}

// QueryHeadersBool 从请求头中获取一个bool类型的值
func QueryHeadersBool(ctx *gin.Context, key string) (has bool, result bool) {
	if gotten, stringResult := QueryHeaders(ctx, key); !gotten {
		return false, false
	} else {
		return convertToBool(stringResult)
	}
}

// QueryHeadersJson 从请求头中获取一个对象
func QueryHeadersJson[T any](ctx *gin.Context, key string) (has bool, result T) {
	if gotten, stringResult := QueryHeaders(ctx, key); !gotten {
		return false, result
	} else {
		return convertToStruct[T](stringResult)
	}
}

// GetClientIP 获取请求客户端的IP
func GetClientIP(ctx *gin.Context) (result string) {
	if addr := ctx.ClientIP(); addr == "::1" {
		return "127.0.0.1"
	} else {
		return addr
	}
}
