package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sunist-c/genius-invokation-simulator-backend/persistence"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http"
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
			middleware.NewQPSLimiter(cfg),
			middleware.NewAuthenticator(cfg),
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

type UploadCardDeckRequest struct {
	Owner           uint     `json:"owner"`
	RequiredPackage []string `json:"required_package"`
	Cards           []uint   `json:"cards"`
	Characters      []uint   `json:"characters"`
}

type UploadCardDeckResponse struct {
	ID              uint     `json:"id"`
	Owner           uint     `json:"owner"`
	RequiredPackage []string `json:"required_package"`
	Cards           []uint   `json:"cards"`
	Characters      []uint   `json:"characters"`
}

type UpdateCardDeckRequest struct {
	Owner           uint     `json:"owner"`
	RequiredPackage []string `json:"required_package"`
	Cards           []uint   `json:"cards"`
	Characters      []uint   `json:"characters"`
}

type UpdateCardDeckResponse struct {
	ID              uint     `json:"id"`
	Owner           uint     `json:"owner"`
	RequiredPackage []string `json:"required_package"`
	Cards           []uint   `json:"cards"`
	Characters      []uint   `json:"characters"`
}

type QueryCardDeckResponse struct {
	ID              uint     `json:"id"`
	Owner           uint     `json:"owner"`
	RequiredPackage []string `json:"required_package"`
	Cards           []uint   `json:"cards"`
	Characters      []uint   `json:"characters"`
}

func uploadDeckServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		request := UploadCardDeckRequest{}
		if !util.BindJson(ctx, &request) {
			// RequestBody解析失败，BadRequest
			ctx.AbortWithStatus(400)
		} else if exist, tokenValue := middleware.GetToken(ctx, cfg); !exist {
			// 服务器找不到token，Forbidden
			ctx.AbortWithStatus(403)
		} else if tokenValue.UID != request.Owner {
			// 上传者和拥有者不同，Forbidden
			ctx.AbortWithStatus(403)
		} else if success, entity := persistence.CardDeckPersistence.InsertOne(&persistence.CardDeck{
			OwnerUID:         request.Owner,
			RequiredPackages: request.RequiredPackage,
			Cards:            request.Cards,
			Characters:       request.Characters,
		}); !success {
			// 插入失败，InternalError
			ctx.AbortWithStatus(500)
		} else {
			// 插入成功，Success
			ctx.JSON(200, UploadCardDeckResponse{
				ID:              entity.ID,
				Owner:           entity.OwnerUID,
				RequiredPackage: entity.RequiredPackages,
				Cards:           entity.Cards,
				Characters:      entity.Characters,
			})
		}
	}
}

func deleteDeckServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if gotten, id := util.QueryPathInt(ctx, ":card_deck_id"); !gotten {
			// 找不到必要的URL路径参数，BadRequest
			ctx.AbortWithStatus(400)
		} else if exist, tokenValue := middleware.GetToken(ctx, cfg); !exist {
			// 服务器找不到token，Forbidden
			ctx.AbortWithStatus(403)
		} else if has, entity := persistence.CardDeckPersistence.QueryByID(uint(id)); !has {
			// 没找到要删除的卡组，NotFound
			ctx.AbortWithStatus(404)
		} else if tokenValue.UID != entity.OwnerUID {
			// 删除者和拥有者不同，Forbidden
			ctx.AbortWithStatus(403)
		} else if success := persistence.CardDeckPersistence.DeleteOne(uint(id)); !success {
			// 删除失败，InternalError
			ctx.AbortWithStatus(500)
		} else {
			// 删除成功，Success
			ctx.Status(200)
		}
	}
}

func updateDeckServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		request := UpdateCardDeckRequest{}
		if !util.BindJson(ctx, &request) {
			// RequestBody解析失败，BadRequest
			ctx.AbortWithStatus(400)
		} else if gotten, id := util.QueryPathInt(ctx, ":card_deck_id"); !gotten {
			// 找不到必要的URL路径参数，BadRequest
			ctx.AbortWithStatus(400)
		} else if exist, tokenValue := middleware.GetToken(ctx, cfg); !exist {
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
			ctx.JSON(200, UpdateCardDeckResponse{
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
			ctx.JSON(200, QueryCardDeckResponse{
				ID:              entity.ID,
				Owner:           entity.OwnerUID,
				RequiredPackage: entity.RequiredPackages,
				Cards:           entity.Cards,
				Characters:      entity.Characters,
			})
		}
	}
}
