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
			// 服务器找不到token，Forbidden
			ctx.AbortWithStatus(403)
		} else if tokenValue.UID != request.Owner {
			// 上传者和拥有者不同，Forbidden
			ctx.AbortWithStatus(403)
		} else if existPlayer, player := persistence.PlayerPersistence.QueryByID(tokenValue.UID); !existPlayer {
			// 没有上传者的玩家记录，理论上不存在这种可能，PreconditionFailed
			ctx.AbortWithStatus(412)
		} else if insertDeckSuccess, cardDeck := persistence.CardDeckPersistence.InsertOne(&persistence.CardDeck{
			OwnerUID:         request.Owner,
			RequiredPackages: request.RequiredPackage,
			Cards:            request.Cards,
			Characters:       request.Characters,
		}); !insertDeckSuccess {
			// 插入记录失败，InternalError
			ctx.AbortWithStatus(500)
		} else {
			// 将新记录添加进玩家信息中
			player.CardDecks = append(player.CardDecks, cardDeck.ID)
			if updatePlayerSuccess := persistence.PlayerPersistence.UpdateByID(request.Owner, player); !updatePlayerSuccess {
				// 更新失败，回滚，InternalError
				persistence.CardDeckPersistence.DeleteOne(cardDeck.ID)
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
}

func deleteDeckServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if gotten, id := util.QueryPathInt(ctx, ":card_deck_id"); !gotten {
			// 找不到必要的URL路径参数，BadRequest
			ctx.AbortWithStatus(400)
		} else if exist, tokenValue := middleware.GetToken(ctx, middlewareConfig); !exist {
			// 服务器找不到token，Forbidden
			ctx.AbortWithStatus(403)
		} else if has, entity := persistence.CardDeckPersistence.QueryByID(uint(id)); !has {
			// 没找到要删除的卡组，NotFound
			ctx.AbortWithStatus(404)
		} else if tokenValue.UID != entity.OwnerUID {
			// 删除者和拥有者不同，Forbidden
			ctx.AbortWithStatus(403)
		} else if existPlayer, player := persistence.PlayerPersistence.QueryByID(tokenValue.UID); !existPlayer {
			// 不存在要删除的玩家，理论上不存在这种可能，PreconditionFailed
			ctx.AbortWithStatus(412)
		} else {
			// 寻找要删除的卡组，并将其移除出玩家持有卡组
			var updatedCardDeck []uint
			for index, cardID := range player.CardDecks {
				if cardID == uint(id) {
					updatedCardDeck = append(player.CardDecks[:index], player.CardDecks[index+1:]...)
					break
				}
			}

			if len(updatedCardDeck) != len(player.CardDecks)-1 {
				// 在玩家的卡组中没有找到要删除的卡组，理论上不存在这种可能，PreconditionFailed
				ctx.AbortWithStatus(412)
			} else if updatePlayerSuccess := persistence.PlayerPersistence.UpdateByID(player.UID, persistence.Player{
				UID:       player.UID,
				NickName:  player.NickName,
				CardDecks: updatedCardDeck,
				Password:  player.Password,
			}); !updatePlayerSuccess {
				// 更新玩家持有卡组失败，InternalError
				ctx.AbortWithStatus(500)
			} else if deleted := persistence.CardDeckPersistence.DeleteOne(uint(id)); !deleted {
				// 删除卡组失败，回滚，InternalError
				persistence.PlayerPersistence.UpdateByID(player.UID, player)
				ctx.AbortWithStatus(500)
			} else {
				// 删除成功，Success
				ctx.Status(200)
			}
		}
	}
}

func updateDeckServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		request := message.UpdateCardDeckRequest{}
		if !util.BindJson(ctx, &request) {
			// RequestBody解析失败，BadRequest
			ctx.AbortWithStatus(400)
		} else if gotten, id := util.QueryPathInt(ctx, ":card_deck_id"); !gotten {
			// 找不到必要的URL路径参数，BadRequest
			ctx.AbortWithStatus(400)
		} else if exist, tokenValue := middleware.GetToken(ctx, middlewareConfig); !exist {
			// 服务器里找不到token，Forbidden
			ctx.AbortWithStatus(403)
		} else if tokenValue.UID != request.Owner {
			// 拥有者和更新者不一致，Forbidden
			ctx.AbortWithStatus(403)
		} else if has, _ := persistence.CardDeckPersistence.QueryByID(uint(id)); !has {
			// 没找到要更新的卡组，NotFound
			ctx.AbortWithStatus(404)
		} else if success := persistence.CardDeckPersistence.UpdateByID(uint(id), persistence.CardDeck{
			ID:               uint(id),
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
				ID:              uint(id),
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
		if gotten, id := util.QueryPathInt(ctx, ":card_deck_id"); !gotten {
			// 找不到必要的URL路径参数，BadRequest
			ctx.AbortWithStatus(400)
		} else if has, entity := persistence.CardDeckPersistence.QueryByID(uint(id)); !has {
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
