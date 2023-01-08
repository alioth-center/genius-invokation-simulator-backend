package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var nullHttpEngine *gin.Engine = nil

var (
	httpEngine *gin.Engine
)

func init() {
	httpEngine = gin.Default()
}

func registerServices() {
	if httpEngine == nullHttpEngine {
		panic(nil)
	}
}

func Serve(port uint, errChan chan error) {
	if err := httpEngine.Run(fmt.Sprintf("0.0.0.0:%v", port)); err != nil {
		errChan <- err
	}
}
