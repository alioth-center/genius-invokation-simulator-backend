package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/middleware"
)

var (
	roomInfoRouter *gin.RouterGroup
)

func initRoomInfoService() {
	roomInfoRouter = http.RegisterServices("/room")

	roomInfoRouter.Use(
		append(
			http.EngineMiddlewares,
			middleware.NewQPSLimiter(middlewareConfig),
		)...,
	)

	roomInfoRouter.GET("",
		listRoomServiceHandler(),
	)
	roomInfoRouter.GET(":room_id",
		queryRoomServiceHandler(),
	)
}

func listRoomServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

	}
}

func queryRoomServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

	}
}
