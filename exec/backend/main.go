package backend

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/service"
)

func Run(port uint) {
	errChan := make(chan error)
	service.InitServices(http.GetConfig().Middleware)
	http.Serve(port, errChan)

	err := <-errChan
	panic(err)
}

func Quit() {}
