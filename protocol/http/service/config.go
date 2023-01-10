package service

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/config"
)

var (
	middlewareConfig config.MiddlewareConfig
	serviceConfig    config.ServiceConfig
)

func SetConfig(conf config.EngineConfig) {
	middlewareConfig = conf.Middleware
	serviceConfig = conf.Service
}
