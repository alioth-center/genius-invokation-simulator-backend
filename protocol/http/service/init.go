package service

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/config"
)

func InitServices(conf config.EngineConfig) {
	SetConfig(conf)
	initPlayerService()
	initLocalizeService()
	initCardDeckService()
}
