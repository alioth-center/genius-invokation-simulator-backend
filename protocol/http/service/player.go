package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sunist-c/genius-invokation-simulator-backend/persistence"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/middleware"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/util"
	util2 "github.com/sunist-c/genius-invokation-simulator-backend/util"
	"time"
)

var (
	playerRouter *gin.RouterGroup
)

func init() {
	playerRouter = http.RegisterServices("/player")
	cfg := http.GetConfig().Middleware
	playerRouter.Use(http.EngineMiddlewares...)
	playerRouter.GET("/login/:player_id",
		middleware.NewQPSLimiter(time.Duration(cfg.QPSLimitTime)*time.Second, cfg.IPTranceKey),
		middleware.NewInterdictor(
			cfg.InterdictorTriggerCount,
			cfg.InterdictorTraceKey,
			time.Duration(cfg.InterdictorBlockedTime)*time.Second,
			cfg.IPTranceKey,
		),
		loginServiceHandler(),
	)
	playerRouter.POST("",
		middleware.NewQPSLimiter(time.Duration(cfg.QPSLimitTime)*time.Second, cfg.IPTranceKey),
		registerServiceHandler(),
	)
	playerRouter.PUT(":player_id/password",
		middleware.NewQPSLimiter(time.Duration(cfg.QPSLimitTime)*time.Second, cfg.IPTranceKey),
		updatePasswordServiceHandler(),
	)
	playerRouter.PUT(":player_id/nickname",
		middleware.NewQPSLimiter(time.Duration(cfg.QPSLimitTime)*time.Second, cfg.IPTranceKey),
		updateNickNameServiceHandler(),
	)
}

type LoginResponse struct {
	PlayerUID       uint                   `json:"player_uid"`
	Success         bool                   `json:"success"`
	PlayerNickName  string                 `json:"player_nick_name"`
	PlayerCardDecks []persistence.CardDeck `json:"player_card_decks"`
}

type LoginRequest struct {
	Password string `json:"password"`
}

type RegisterRequest struct {
	NickName string `json:"nick_name"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	PlayerUID      uint   `json:"player_uid"`
	PlayerNickName string `json:"player_nick_name"`
}

type UpdatePasswordRequest struct {
	OriginalPassword string `json:"original_password"`
	NewPassword      string `json:"new_password"`
}

type UpdateNickNameRequest struct {
	Password    string `json:"password"`
	NewNickName string `json:"new_nick_name"`
}

func loginServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		request := LoginRequest{}
		if !util.BindJson(ctx, &request) {
			// RequestBody解析失败，BadRequest
			ctx.AbortWithStatus(400)
		} else if gotten, id := util.QueryPathInt(ctx, ":player_id"); !gotten {
			// 找不到必要的URL路径参数，BadRequest
			ctx.AbortWithStatus(400)
		} else if has, player := persistence.PlayerPersistence.QueryByID(uint(id)); !has {
			// 没找到请求玩家，NotFound，登陆失败
			ctx.JSON(404, LoginResponse{Success: false})
		} else if success, encodeResult := util2.EncodePassword([]byte(request.Password), uint(id)); !success {
			// 编码密码失败，InternalError，登陆失败
			ctx.JSON(500, LoginResponse{Success: false})
		} else if string(encodeResult) != (player.Password) {
			// 密码校验失败，Forbidden，登陆失败
			ctx.JSON(403, LoginResponse{Success: false})
		} else {
			// 登录成功，获取玩家卡组信息后返回登录成功响应
			response := LoginResponse{
				PlayerUID:       player.UID,
				Success:         true,
				PlayerNickName:  player.NickName,
				PlayerCardDecks: []persistence.CardDeck{},
			}
			for _, cardDeckID := range player.CardDecks {
				if success, cardDeck := persistence.CardDeckPersistence.QueryByID(cardDeckID); success {
					response.PlayerCardDecks = append(response.PlayerCardDecks, cardDeck)
				}
			}
			ctx.JSON(200, response)
		}
	}
}

func registerServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		request := RegisterRequest{}
		if !util.BindJson(ctx, &request) {
			// RequestBody解析失败，BadRequest
			ctx.AbortWithStatus(400)
		} else if success, id := persistence.PlayerPersistence.InsertOne(persistence.Player{NickName: request.NickName}); !success {
			// 创建Player失败，InternalError
			ctx.AbortWithStatus(500)
		} else if encoded, encodedPassword := util2.EncodePassword([]byte(request.Password), id); !encoded {
			// 编码密码失败，回滚，InternalError
			persistence.PlayerPersistence.DeleteOne(id)
			ctx.AbortWithStatus(500)
		} else if updated := persistence.PlayerPersistence.UpdateByID(id, persistence.Player{UID: id, NickName: request.NickName, Password: string(encodedPassword)}); !updated {
			// 更新密码失败，回滚，InternalError
			persistence.PlayerPersistence.DeleteOne(id)
			ctx.AbortWithStatus(500)
		} else {
			// 注册Player成功，返回Player信息
			ctx.JSON(200, RegisterResponse{
				PlayerUID:      id,
				PlayerNickName: request.NickName,
			})
		}
	}
}

func updatePasswordServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
	}
}

func updateNickNameServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

	}
}

func destroyServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

	}
}
