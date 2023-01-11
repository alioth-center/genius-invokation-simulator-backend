package config

var (
	websocketConfig *WebsocketConfig
)

type WebsocketConfig struct {
	HandshakeTimeout uint `json:"handshake_timeout" yaml:"handshake_timeout" xml:"handshake_timeout"`
}

func SetConfig(conf WebsocketConfig) {
	*websocketConfig = conf
}

func GetConfig() WebsocketConfig {
	return *websocketConfig
}
