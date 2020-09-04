package web

import (
	"bufio"
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"time"
)

//session列表
var Sessions = make(map[string]*WebSocketConn)

type WebSocketConn struct {
	*websocket.Conn
	ID         string
	createTime time.Time
	ctx        context.Context
	cancel     context.CancelFunc
}

func (c *WebSocketConn) Start() {
	//webview主动关闭
	c.SetCloseHandler(func(code int, text string) error {
		c.Destroy()
		return nil
	})
	Sessions[c.ID] = c
	//读取数据错误,关闭措施
	defer c.Destroy()

	go c.StartReader()
	go c.StartWriter()
	<-c.ctx.Done()
}

func (c *WebSocketConn) Destroy() {
	//fmt.Printf("close session:%s \n", c.ID)
	c.cancel()
	if c.Conn != nil {
		_ = c.Conn.Close()
		c.Conn = nil
		delete(Sessions, c.ID)
	}
}

func (c *WebSocketConn) StartReader() {
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			msgType, r, err := c.NextReader()
			if err != nil {
				print(err.Error())
				goto result
			}
			if msgType != websocket.TextMessage {
				err := fmt.Errorf("[web]read message not support messageType:%v", msgType)
				fmt.Println(err)
				print(err.Error())
				goto result
			}

			reader := bufio.NewReader(r)
			line, _, err := reader.ReadLine()
			if err != nil && err == io.EOF {
				println("关闭")

				goto result
			}
			if err != nil || err != io.EOF {
				print(err.Error())
				goto result
			}
			//读取到数据后,发送到shell
			fmt.Println(line)
		}
	}
result:
	c.cancel()
}

func (c *WebSocketConn) StartWriter() {
	err := c.WriteMessage(websocket.TextMessage, []byte(c.ID))
	if err != nil {
		print(err.Error())
		goto result
	}
	//从shell中读取数据,并发送回前端
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			err := c.WriteMessage(websocket.TextMessage, []byte(time.Now().Format("2006-01-02 15:04:05")))
			if err != nil {
				print(err.Error())
				goto result
			}
			time.Sleep(time.Second)
		}
	}
result:
	c.cancel()
}

func NewWebSocketConn(ctx context.Context, conn *websocket.Conn, ID string) *WebSocketConn {
	ctx, cancelFunc := context.WithCancel(ctx)
	return &WebSocketConn{Conn: conn, ID: ID, createTime: time.Now(), ctx: ctx, cancel: cancelFunc}
}
