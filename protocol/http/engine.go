package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/middleware"
)

var nullHttpEngine *gin.Engine = nil

var (
	httpEngine *gin.Engine
)

var (
	EngineMiddlewares = []gin.HandlerFunc{
		middleware.NewIPTracer(config.Middleware.UUIDKey),   // IP追踪器
		middleware.NewUUIDTagger(config.Middleware.UUIDKey), // UUID标记器
	}
)

func init() {
	httpEngine = gin.Default()
}

func RegisterServices(subPath string) *gin.RouterGroup {
	if httpEngine == nullHttpEngine {
		panic("nil http engine")
	} else {
		return httpEngine.Group(subPath)
	}
}

func Serve(port uint, errChan chan error) {
	if err := httpEngine.Run(fmt.Sprintf("0.0.0.0:%v", port)); err != nil {
		errChan <- err
	}
}
