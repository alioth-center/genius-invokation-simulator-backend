package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sunist-c/genius-invokation-simulator-backend/persistence"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/middleware"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/util"
)

var (
	roomInfoRouter *gin.RouterGroup
	roomInfoCache  map[uint64]bool
)

func initRoomInfoService() {
	roomInfoRouter = http.RegisterServices("/room")
	roomInfoCache = map[uint64]bool{}
	for i := uint64(0); i < serviceConfig.MaxRooms; i++ {
		roomInfoCache[i] = false
	}

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

type ListRoomResponse struct {
	Rooms []persistence.RoomInfo `json:"rooms"`
}

type QueryRoomResponse struct {
	RoomInfo persistence.RoomInfo `json:"room_info"`
}

func listRoomServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		response := &ListRoomResponse{Rooms: []persistence.RoomInfo{}}

		// 查询RoomInfo状态
		for roomID, roomValid := range roomInfoCache {
			// 将可用的房间进行查询
			if roomValid {
				if success, appendRoom := persistence.RoomInfoPersistence.QueryByID(roomID); !success {
					// 查找RoomInfo失败，InternalError
					ctx.AbortWithStatus(500)
					return
				} else {
					// 查找RoomInfo成功，加入到响应中
					response.Rooms = append(response.Rooms, appendRoom)
				}
			}
		}

		// 查询完毕，Success
		ctx.JSON(200, response)
	}
}

func queryRoomServiceHandler() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if existRoomID, roomID := util.QueryPathUint64(ctx, ":room_id"); !existRoomID {
			// 缺少必要的URL参数，BadRequest
			ctx.AbortWithStatus(400)
		} else if roomValid, existRoom := roomInfoCache[roomID]; !existRoom || !roomValid {
			// 请求的房间不存在或无效，NotFound
			ctx.AbortWithStatus(404)
		} else if existRoomInfo, roomInfo := persistence.RoomInfoPersistence.QueryByID(roomID); !existRoomInfo {
			// 查询房间信息失败，InternalError
			ctx.AbortWithStatus(500)
		} else {
			// 查询成功，Success
			response := QueryRoomResponse{RoomInfo: roomInfo}
			ctx.JSON(200, response)
		}
	}
}
