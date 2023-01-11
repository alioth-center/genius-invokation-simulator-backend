package config

var (
	config = &EngineConfig{
		Middleware: MiddlewareConfig{
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
		Service: ServiceConfig{
			MaxRooms: 100,
		},
	}
)

type EngineConfig struct {
	Middleware MiddlewareConfig `json:"middleware" yaml:"middleware" xml:"middleware"`
	Service    ServiceConfig    `json:"service" yaml:"service" xml:"service"`
}

func SetConfig(conf EngineConfig) {
	config = &conf
}

func GetConfig() EngineConfig {
	return *config
}
