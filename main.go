package main

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/protocal/http"
	"time"
)

func main() {
	http.StartHttpServer(8080)
	time.Sleep(10 * time.Second)
}
