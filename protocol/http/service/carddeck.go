package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sunist-c/genius-invokation-simulator-backend/persistence"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/message"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/middleware"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/util"
)

var (
	deckRouter *gin.RouterGroup
)

func initCardDeckService() {
	deckRouter = http.RegisterServices("/card_deck")

	deckRouter.Use(
		append(
			http.EngineMiddlewares,
			middleware.NewQPSLimiter(middlewareConfig),
			middleware.NewAuthenticator(middlewareConfig),
		)...,
	)

	deckRouter.POST("",
		uploadDeckServiceHandler(),
	)
	deckRouter.GET(":card_deck_id",
		queryDeckServiceHandler(),
	)
	deckRouter.PUT(":card_deck_id",
		updateDeckServiceHandler(),
	)
	deckRouter.DELETE(":card_deck_id",
		deleteDeckServiceHandler(),
	)
}

func uploadDeckServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		request := message.UploadCardDeckRequest{}
		if !util.BindJson(ctx, &request) {
			// RequestBody解析失败，BadRequest
			ctx.AbortWithStatus(400)
		} else if exist, tokenValue := middleware.GetToken(ctx, middlewareConfig); !exist {
			// 服务器找不到token，UnAuthorized
			ctx.AbortWithStatus(401)
		} else if tokenValue.UID != request.Owner {
			// 上传者和已验证者不同，UnAuthorized
			ctx.AbortWithStatus(401)
		} else if existPlayer, _ := persistence.PlayerPersistence.QueryByID(tokenValue.UID); !existPlayer {
			// 没有上传者的玩家记录，理论上不存在这种可能，NotFound
			ctx.AbortWithStatus(404)
		} else if insertDeckSuccess, cardDeck := persistence.CardDeckPersistence.InsertOne(&persistence.CardDeck{
			OwnerUID:         request.Owner,
			RequiredPackages: request.RequiredPackage,
			Cards:            request.Cards,
			Characters:       request.Characters,
		}); !insertDeckSuccess {
			// 插入记录失败，InternalError
			ctx.AbortWithStatus(500)
		} else {
			// 更新成功，Success
			ctx.JSON(200, message.UploadCardDeckResponse{
				ID:              cardDeck.ID,
				Owner:           cardDeck.OwnerUID,
				RequiredPackage: cardDeck.RequiredPackages,
				Cards:           cardDeck.Cards,
				Characters:      cardDeck.Characters,
			})
		}
	}
}

func deleteDeckServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if gotten, id := util.QueryPathUint64(ctx, ":card_deck_id"); !gotten {
			// 找不到必要的URL路径参数，BadRequest
			ctx.AbortWithStatus(400)
		} else if exist, tokenValue := middleware.GetToken(ctx, middlewareConfig); !exist {
			// 服务器找不到token，UnAuthorized
			ctx.AbortWithStatus(401)
		} else if has, entity := persistence.CardDeckPersistence.QueryByID(id); !has {
			// 没找到要删除的卡组，NotFound
			ctx.AbortWithStatus(404)
		} else if tokenValue.UID != entity.OwnerUID {
			// 删除者和拥有者不同，Forbidden
			ctx.AbortWithStatus(403)
		} else if existPlayer, _ := persistence.PlayerPersistence.QueryByID(tokenValue.UID); !existPlayer {
			// 不存在要删除的玩家，理论上不存在这种可能，NotFound
			ctx.AbortWithStatus(404)
		} else if deleted := persistence.CardDeckPersistence.DeleteOne(id); !deleted {
			// 删除卡组失败，InternalError
			ctx.AbortWithStatus(500)
		} else {
			// 删除成功，Success
			ctx.Status(200)
		}

	}
}

func updateDeckServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		request := message.UpdateCardDeckRequest{}
		if !util.BindJson(ctx, &request) {
			// RequestBody解析失败，BadRequest
			ctx.AbortWithStatus(400)
		} else if gotten, id := util.QueryPathUint64(ctx, ":card_deck_id"); !gotten {
			// 找不到必要的URL路径参数，BadRequest
			ctx.AbortWithStatus(400)
		} else if exist, tokenValue := middleware.GetToken(ctx, middlewareConfig); !exist {
			// 服务器里找不到token，UnAuthorized
			ctx.AbortWithStatus(401)
		} else if tokenValue.UID != request.Owner {
			// 拥有者和更新者不一致，Forbidden
			ctx.AbortWithStatus(403)
		} else if has, _ := persistence.CardDeckPersistence.QueryByID(id); !has {
			// 没找到要更新的卡组，NotFound
			ctx.AbortWithStatus(404)
		} else if success := persistence.CardDeckPersistence.UpdateByID(
			id,
			persistence.CardDeck{
				ID:               id,
				OwnerUID:         request.Owner,
				RequiredPackages: request.RequiredPackage,
				Cards:            request.Cards,
				Characters:       request.Characters,
			}); !success {
			// 更新失败，InternalError
			ctx.AbortWithStatus(500)
		} else {
			// 更新成功，Success
			ctx.JSON(200, message.UpdateCardDeckResponse{
				ID:              id,
				Owner:           request.Owner,
				RequiredPackage: request.RequiredPackage,
				Cards:           request.Cards,
				Characters:      request.Characters,
			})
		}
	}
}

func queryDeckServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if gotten, id := util.QueryPathUint64(ctx, ":card_deck_id"); !gotten {
			// 找不到必要的URL路径参数，BadRequest
			ctx.AbortWithStatus(400)
		} else if has, entity := persistence.CardDeckPersistence.QueryByID(id); !has {
			// 找不到卡组，NotFound
			ctx.AbortWithStatus(404)
		} else if has {
			// 找到了卡组，Success
			ctx.JSON(200, message.QueryCardDeckResponse{
				ID:              entity.ID,
				Owner:           entity.OwnerUID,
				RequiredPackage: entity.RequiredPackages,
				Cards:           entity.Cards,
				Characters:      entity.Characters,
			})
		}
	}
}
