package config

var (
	TcpConfig = &TcpConfig{
		port:    "8888",
		ip_addr: "127.0.0.1",
		bufsize: 4096,
	}
)

type TcpConfig struct{
	port string
	ip_addr string
	bufsize int
}

func SetConfig(conf TcpConfig) {
	*tcpConfig = conf
}

func GetConfig() WebsocketConfig {
	return *tcpConfig
}