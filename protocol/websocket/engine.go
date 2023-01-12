package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sunist-c/genius-invokation-simulator-backend/protocol/websocket/config"
	"net/http"
	"time"
)

func checkOriginHandler() func(r *http.Request) bool {
	conf := config.GetConfig()
	if conf.AllowCrossOrigin {
		return func(r *http.Request) bool {
			return true
		}
	} else {
		return func(r *http.Request) bool {
			// todo: optimize check origin logic
			origin := r.Header.Get("Origin")
			for _, allow := range conf.AllowOriginDomains {
				if origin == allow {
					return true
				}
			}

			return false
		}
	}
}

func newUpgrader() *websocket.Upgrader {
	conf := config.GetConfig()
	return &websocket.Upgrader{
		HandshakeTimeout: time.Second * time.Duration(conf.HandshakeTimeout),
		ReadBufferSize:   int(conf.WebsocketReaderBufferSize),
		WriteBufferSize:  int(conf.WebsocketWriterBufferSize),
		CheckOrigin:      checkOriginHandler(),
	}
}

type Connection struct {
	conn     *websocket.Conn
	exitChan chan struct{}
	errChan  chan error
	iStream  chan []byte // iStream 输入管道
	oStream  chan []byte // oStream 输出管道
}

func (c *Connection) readLoop() {
	for {
		select {
		case <-c.exitChan:
			// 向exitChan写入信息通知其他协程
			c.exitChan <- struct{}{}
			return
		default:
			if messageType, reader, err := c.conn.ReadMessage(); err != nil {
				// 发生错误，关闭WebSocket连接
				c.errChan <- err
				c.Close()
			} else if messageType == websocket.TextMessage {
				// 传输格式为json，只监听TextMessage
				c.iStream <- reader
			}
		}
	}
}

func (c *Connection) writeLoop() {
	for {
		select {
		case <-c.exitChan:
			// 向exitChan写入信息通知其他协程
			c.exitChan <- struct{}{}
			return
		case outMessage := <-c.oStream:
			if err := c.conn.WriteMessage(websocket.TextMessage, outMessage); err != nil {
				// 发生错误，关闭WebSocket连接
				c.errChan <- err
				c.Close()
			}
		}
	}
}

// Serve 启动WebsocketConnection的服务协程
func (c *Connection) Serve() {
	go c.readLoop()
	go c.writeLoop()
}

func (c *Connection) Close() {
	// 关闭客户端连接
	if err := c.conn.Close(); err != nil {
		c.errChan <- err
	}

	// 通知readLoop和writeLoop退出
	c.exitChan <- struct{}{}

	// 释放WebSocketConnection
	c.conn = nil
	close(c.iStream)
	close(c.oStream)
	close(c.exitChan)
}

func (c *Connection) Write(jsonBytes []byte) {
	c.oStream <- jsonBytes
}

func (c *Connection) Read() <-chan []byte {
	return c.iStream
}

func NewConnection(conn *websocket.Conn, errChan chan error) *Connection {
	conf := config.GetConfig()
	return &Connection{
		conn:     conn,
		iStream:  make(chan []byte, conf.ServerMessageBufferSize),
		oStream:  make(chan []byte, conf.ServerMessageBufferSize),
		exitChan: make(chan struct{}, 2),
		errChan:  errChan,
	}
}

// Upgrade 将HTTP连接升级为Websocket连接
func Upgrade(ctx *gin.Context, errChan chan error) (success bool, connection *Connection) {
	if conn, err := newUpgrader().Upgrade(ctx.Writer, ctx.Request, ctx.Request.Header); err != nil {
		return false, nil
	} else {
		return true, NewConnection(conn, errChan)
	}
}
