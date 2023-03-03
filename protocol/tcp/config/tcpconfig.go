package config

var (
	TcpConfig = &TcpConfig{
		port:    "8888",
		ip_addr: "127.0.0.1",
	}
)

type TcpConfig struct{
	port string
	ip_addr string
}

func SetConfig(conf TcpConfig) {
	*tcpConfig = conf
}

func GetConfig() TcpConfig {
	return *tcpConfig
}