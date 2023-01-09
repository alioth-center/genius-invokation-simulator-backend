package service

import "github.com/sunist-c/genius-invokation-simulator-backend/protocol/http"

var (
	cfg http.MiddlewareConfig
)

func InitServices(conf http.MiddlewareConfig) {
	cfg = conf
	initializePlayerService()
	initLocalizeService()
}
