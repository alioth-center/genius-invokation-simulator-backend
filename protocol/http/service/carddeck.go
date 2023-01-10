package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/middleware"
	"time"
)

var (
	cardDeckRouter *gin.RouterGroup
)

func initCardDeckService() {
	cardDeckRouter = http.RegisterServices("/card_deck")

	cardDeckRouter.Use(
		append(
			http.EngineMiddlewares,
			middleware.NewQPSLimiter(time.Duration(cfg.QPSLimitTime)*time.Second, cfg.IPTranceKey),
		)...,
	)
}
