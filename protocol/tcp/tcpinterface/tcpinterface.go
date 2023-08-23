package tcpinterface

import "net"

type TcpConnection interface{
	Upgrade(ctx *gin.Context) net.Conn

	Read() chan []byte

	Write(jsonBytes []byte)
}