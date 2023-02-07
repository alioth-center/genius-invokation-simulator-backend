package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sunist-c/genius-invokation-simulator-backend/persistence"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/message"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/middleware"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/util"
	util2 "github.com/sunist-c/genius-invokation-simulator-backend/util"
)

var (
	playerRouter *gin.RouterGroup
)

func initPlayerService() {
	playerRouter = http.RegisterServices("/player")

	playerRouter.Use(
		append(
			http.EngineMiddlewares,
			middleware.NewQPSLimiter(middlewareConfig),
		)...,
	)

	playerRouter.POST("/login",
		middleware.NewInterdictor(middlewareConfig),
		loginServiceHandler(),
	)
	playerRouter.POST("",
		registerServiceHandler(),
	)
	playerRouter.PATCH(":player_id/password",
		middleware.NewInterdictor(middlewareConfig),
		updatePasswordServiceHandler(),
	)
	playerRouter.PATCH(":player_id/nickname",
		middleware.NewInterdictor(middlewareConfig),
		updateNickNameServiceHandler(),
	)
	playerRouter.DELETE(":player_id",
		middleware.NewInterdictor(middlewareConfig),
		destroyServiceHandler(),
	)
}

func loginServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		request := message.LoginRequest{}
		if !util.BindJson(ctx, &request) {
			// RequestBody解析失败，BadRequest
			ctx.AbortWithStatus(400)
		} else if has, player := persistence.PlayerPersistence.QueryByID(request.PlayerUID); !has {
			// 没找到请求玩家，NotFound，登陆失败
			ctx.JSON(404, message.LoginResponse{Success: false})
		} else if success, encodeResult := util2.EncodePassword([]byte(request.Password), uint(request.PlayerUID)); !success {
			// 编码密码失败，InternalError，登陆失败
			ctx.JSON(500, message.LoginResponse{Success: false})
		} else if string(encodeResult) != (player.Password) {
			// 密码校验失败，Forbidden，登陆失败
			middleware.Interdict(ctx, middlewareConfig)
			ctx.JSON(403, message.LoginResponse{Success: false})
		} else if !middleware.AttachToken(ctx, middlewareConfig, request.PlayerUID) {
			// 生成token失败，InternalError
			ctx.JSON(500, message.LoginResponse{Success: false})
		} else {
			// 登录成功，获取玩家卡组信息后返回登录成功响应
			wantDeck := persistence.CardDeck{OwnerUID: player.UID}
			resultDeck := persistence.CardDeckPersistence.Find(wantDeck)
			response := message.LoginResponse{
				PlayerUID:       player.UID,
				Success:         true,
				PlayerNickName:  player.NickName,
				PlayerCardDecks: make([]message.CardDeck, len(resultDeck)),
			}

			for i, deck := range resultDeck {
				response.PlayerCardDecks[i] = message.CardDeck{
					ID:               deck.ID,
					OwnerUID:         deck.OwnerUID,
					RequiredPackages: deck.RequiredPackages,
					Cards:            deck.Cards,
					Characters:       deck.Characters,
				}
			}

			ctx.JSON(200, response)
		}
	}
}

func registerServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		request := message.RegisterRequest{}
		if !util.BindJson(ctx, &request) {
			// RequestBody解析失败，BadRequest
			ctx.AbortWithStatus(400)
		} else if success, result := persistence.PlayerPersistence.InsertOne(
			&persistence.Player{
				NickName: request.NickName,
			}); !success {
			// 创建Player失败，InternalError
			ctx.AbortWithStatus(500)
		} else if encoded, encodedPassword := util2.EncodePassword([]byte(request.Password), result.UID); !encoded {
			// 编码密码失败，回滚，InternalError
			persistence.PlayerPersistence.DeleteOne(result.UID)
			ctx.AbortWithStatus(500)
		} else if updated := persistence.PlayerPersistence.UpdateByID(result.UID,
			persistence.Player{
				UID:      result.UID,
				NickName: request.NickName,
				Password: string(encodedPassword),
			}); !updated {
			// 更新密码失败，回滚，InternalError
			persistence.PlayerPersistence.DeleteOne(result.UID)
			ctx.AbortWithStatus(500)
		} else {
			// 注册Player成功，返回Player信息
			ctx.JSON(200, message.RegisterResponse{
				PlayerUID:      result.UID,
				PlayerNickName: request.NickName,
			})
		}
	}
}

func updatePasswordServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		request := message.UpdatePasswordRequest{}
		if !util.BindJson(ctx, &request) {
			// RequestBody解析失败，BadRequest
			ctx.AbortWithStatus(400)
		} else if hasID, id := util.QueryPathUint64(ctx, ":player_id"); !hasID {
			// 没有必须的player_id字段，BadRequest
			ctx.AbortWithStatus(400)
		} else if success, player := persistence.PlayerPersistence.QueryByID(id); !success {
			// 没有找到玩家，NotFound
			ctx.AbortWithStatus(404)
		} else if encoded, encodedPassword := util2.EncodePassword([]byte(request.OriginalPassword), id); !encoded {
			// 编码原密码失败，InternalError
			ctx.AbortWithStatus(500)
		} else if string(encodedPassword) != player.Password {
			// 提供的原密码密码不匹配，失败，Forbidden
			middleware.Interdict(ctx, middlewareConfig)
			ctx.AbortWithStatus(403)
		} else if encodedNew, encodedNewPassword := util2.EncodePassword([]byte(request.NewPassword), id); !encodedNew {
			// 编码新密码失败，InternalError
			ctx.AbortWithStatus(500)
		} else if updated := persistence.PlayerPersistence.UpdateByID(
			id,
			persistence.Player{
				UID:      player.UID,
				NickName: player.NickName,
				Password: string(encodedNewPassword),
			}); !updated {
			// 更新新密码失败，InternalError
			ctx.AbortWithStatus(500)
		} else {
			// 更新密码成功，Success
			ctx.AbortWithStatus(200)
		}
	}
}

func updateNickNameServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		request := message.UpdateNickNameRequest{}
		if !util.BindJson(ctx, &request) {
			// RequestBody解析失败，BadRequest
			ctx.AbortWithStatus(400)
		} else if hasID, id := util.QueryPathUint64(ctx, ":player_id"); !hasID {
			// 没有必须的player_id字段，BadRequest
			ctx.AbortWithStatus(400)
		} else if success, player := persistence.PlayerPersistence.QueryByID(id); !success {
			// 没有找到玩家，NotFound
			ctx.AbortWithStatus(404)
		} else if encoded, encodedPassword := util2.EncodePassword([]byte(request.Password), id); !encoded {
			// 编码原密码失败，InternalError
			ctx.AbortWithStatus(500)
		} else if string(encodedPassword) != player.Password {
			// 提供的原密码密码不匹配，失败，Forbidden
			middleware.Interdict(ctx, middlewareConfig)
			ctx.AbortWithStatus(403)
		} else if updated := persistence.PlayerPersistence.UpdateByID(
			id,
			persistence.Player{
				UID:      player.UID,
				NickName: request.NewNickName,
				Password: player.Password,
			},
		); !updated {
			// 更新新昵称失败，InternalError
			ctx.AbortWithStatus(500)
		} else {
			// 更新昵称成功，Success
			ctx.AbortWithStatus(200)
		}
	}
}

func destroyServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		request := message.DestroyPlayerRequest{}
		if !util.BindJson(ctx, &request) {
			// RequestBody解析失败，BadRequest
			ctx.AbortWithStatus(400)
		} else if hasID, id := util.QueryPathUint64(ctx, ":player_id"); !hasID {
			// 没有必须的player_id字段，BadRequest
			ctx.AbortWithStatus(400)
		} else if !request.Confirm {
			// 确认失败，Success
			ctx.JSON(200, message.DestroyPlayerResponse{
				Success: false,
			})
		} else if success, player := persistence.PlayerPersistence.QueryByID(id); !success {
			// 没有找到玩家，NotFound
			ctx.AbortWithStatus(404)
		} else if encoded, encodedPassword := util2.EncodePassword([]byte(request.Password), id); !encoded {
			// 编码原密码失败，InternalError
			ctx.AbortWithStatus(500)
		} else if string(encodedPassword) != player.Password {
			// 提供的原密码密码不匹配，失败，Forbidden
			middleware.Interdict(ctx, middlewareConfig)
			ctx.AbortWithStatus(403)
		} else if destroyed := persistence.PlayerPersistence.DeleteOne(id); !destroyed {
			// 删除失败，InternalError
			ctx.AbortWithStatus(500)
		} else {
			// 删除成功，Success
			ctx.JSON(200, message.DestroyPlayerResponse{
				Success: true,
			})
		}
	}
}
