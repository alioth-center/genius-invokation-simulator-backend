package service

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/middleware"
)

var (
	cfg middleware.Config
)

func InitServices(conf middleware.Config) {
	cfg = conf
	initPlayerService()
	initLocalizeService()
	initCardDeckService()
}
