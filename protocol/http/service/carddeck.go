package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/middleware"
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
		)...,
	)

	deckRouter.POST("",
		middleware.NewQPSLimiter(cfg),
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

func uploadDeckServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

	}
}

func deleteDeckServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

	}
}

func updateDeckServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

	}
}

func queryDeckServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

	}
}
