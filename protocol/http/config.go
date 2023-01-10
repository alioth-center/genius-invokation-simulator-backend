package http

import "github.com/sunist-c/genius-invokation-simulator-backend/protocol/http/middleware"

var (
	config = &Config{
		Middleware: middleware.Config{
			UUIDKey:                 "uuid",
			IPTranceKey:             "ip",
			InterdictorTraceKey:     "interdicted",
			InterdictorBlockedTime:  600,
			InterdictorTriggerCount: 5,
			QPSLimitTime:            1,
			TokenIDKey:              "gisb_token_id",
			TokenKey:                "gisb_token",
			TokenRefreshTime:        7200,
			TokenRemainingTime:      1800,
			CookieDomain:            "localhost",
		},
		Backend: BackendConfig{
			ListenPort: 8086,
		},
	}
)

type Config struct {
	Middleware middleware.Config `json:"middleware"`
	Backend    BackendConfig     `json:"backend"`
}

func SetConfig(conf Config) {
	config = &conf
}

func GetConfig() Config {
	return *config
}

type BackendConfig struct {
	ListenPort uint `json:"listen_port"`
}
