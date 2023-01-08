package backend

import "github.com/sunist-c/genius-invokation-simulator-backend/protocol/http"

func Run(port uint) {
	errChan := make(chan error)
	http.Serve(port, errChan)

	err := <-errChan
	panic(err)
}

func Quit() {}
