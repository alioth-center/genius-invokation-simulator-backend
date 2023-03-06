package tcp

import(
	"fmt"
	"net"
	"github.com/SWUdaishipeng/genius-invokation-simulator-backend/tree/master/protocol/tcp/tcpinterface"
	"github.com/gin-gonic/gin"
)

func (c TcpConnection)Upgrade(ctx *gin.Context,port string) net.Conn{
	ip_addr:=ctx.Request.Header.Get("X-Forwarded-For") 
	listen,err := net.Listen("tcp",ip_addr+":"+port)
	if err != nil {
		fmt.Println("start error,err:",err)
		return
	}
	defer listen.Close()
    
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
	ch <- buf[:n]
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