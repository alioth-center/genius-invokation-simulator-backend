package config

var (
	websocketConfig = &WebsocketConfig{
		HandshakeTimeout:          60,
		AllowCrossOrigin:          true,
		WebsocketWriterBufferSize: 4096,
		WebsocketReaderBufferSize: 4096,
		ServerMessageBufferSize:   128,
		AllowOriginDomains:        []string{""},
	}
)

type WebsocketConfig struct {
	HandshakeTimeout          uint     `json:"handshake_timeout" yaml:"handshake_timeout" xml:"handshake_timeout"`
	AllowCrossOrigin          bool     `json:"allow_cross_origin" yaml:"allow_cross_origin" xml:"allow_cross_origin"`
	WebsocketWriterBufferSize uint     `json:"websocket_writer_buffer_size" yaml:"websocket_writer_buffer_size" xml:"websocket_writer_buffer_size"`
	WebsocketReaderBufferSize uint     `json:"websocket_reader_buffer_size" yaml:"websocket_reader_buffer_size" xml:"websocket_reader_buffer_size"`
	ServerMessageBufferSize   uint     `json:"server_message_buffer_size" yaml:"server_message_buffer_size" xml:"server_message_buffer_size"`
	AllowOriginDomains        []string `json:"allow_origin_domains" yaml:"allow_origin_domains" xml:"allow_origin_domains"`
}

func SetConfig(conf WebsocketConfig) {
	*websocketConfig = conf
}

func GetConfig() WebsocketConfig {
	return *websocketConfig
}
