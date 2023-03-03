package tcp

import(
	"fmt"
	"net"
	"./tcpinterface"
	"./config"
)

func (c TcpConnection)Upgrade(ctx *gin.Context) net.Conn{
	listen,err := net.Listen("tcp",TcpConfig.GetConfig().ip_addr+":"+TcpConfig.GetConfig().port)
	if err != nil {
		fmt.Println("start error,err:",err)
		return
	}
	defer listene.Close()
    
	for {
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("accept error,err:",err)
			continue	// wait client
		}
	return conn
	}
}

func (c TcpConnection)Read(conn net.Conn) chan []byte{
	defer conn.Close()
	var buf [1024]byte
	n,err := conn.Read(buf[:])
	if err != nil {
		fmt.Println("read err:",err)
		return
	}
	ch := make(chan []byte)
	ch<- buf[:n]
	return ch
}

func (c TcpConnection)Write(jsonBytes []byte,conn net.Conn){
	defer conn.Close()
	_,err = conn.Write(jsonBytes)
	if err != nil {
		fmt.Println("send message error,err:",err)
		return
	}
}