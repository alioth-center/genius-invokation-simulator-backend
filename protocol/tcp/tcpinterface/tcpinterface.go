package tcpinterface

import "net"

type TcpConnection interface{
	Upgrade()

	Read() chan []byte

	Write(jsonBytes []byte)
}