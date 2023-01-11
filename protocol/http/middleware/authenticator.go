package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sunist-c/genius-invokation-simulator-backend/persistence"
	"github.com/sunist-c/genius-invokation-simulator-backend/util"
	"strconv"
	"time"
)

// AttachToken 将Token信息附加给响应
func AttachToken(ctx *gin.Context, conf Config, player uint) (success bool) {
	uuid := GetUUID(ctx, conf)
	ok, ip := GetIPTrace(ctx, conf)
	if !ok {
		return false
	}
	tokenID, token := util.GenerateMD5(strconv.Itoa(int(ip))), util.GenerateMD5(uuid)
	if success, _ = persistence.TokenPersistence.InsertOne(token, persistence.Token{UID: player, ID: tokenID}, time.Second*time.Duration(conf.TokenRefreshTime)); !success {
		return false
	} else {
		ctx.SetCookie(conf.TokenIDKey, tokenID, int(conf.TokenRefreshTime), "/", conf.CookieDomain, false, true)
		ctx.SetCookie(conf.TokenKey, token, int(conf.TokenRefreshTime), "/", conf.CookieDomain, false, true)
		return true
	}
}

// GetToken 获取Cookie里的Token信息
func GetToken(ctx *gin.Context, conf Config) (success bool, token persistence.Token) {
	if result, err := ctx.Cookie(conf.TokenKey); err != nil {
		return false, token
	} else if has, tokenStruct, _ := persistence.TokenPersistence.QueryByID(result); !has {
		return false, token
	} else {
		return true, tokenStruct
	}
}

// NewAuthenticator 新建一个认证器，只有Cookie中的Token和ID正确时才会放行，但是没有具体处理是否有权限
func NewAuthenticator(conf Config) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if tokenID, err := ctx.Cookie(conf.TokenIDKey); err != nil {
			ctx.AbortWithStatus(403)
		} else if token, err := ctx.Cookie(conf.TokenKey); err != nil {
			ctx.AbortWithStatus(403)
		} else if has, result, timeoutAt := persistence.TokenPersistence.QueryByID(token); !has {
			ctx.AbortWithStatus(403)
		} else if result.ID != tokenID {
			ctx.AbortWithStatus(403)
		} else if time.Now().Add(time.Duration(conf.TokenRemainingTime) * time.Second).After(timeoutAt) {
			persistence.TokenPersistence.RefreshByID(token, time.Second*time.Duration(conf.TokenRefreshTime))
		}
	}
}
