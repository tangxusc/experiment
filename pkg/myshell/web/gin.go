package web

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
	"testgo/pkg/myshell/config"
)

const IndexPath = "/index"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func StartWebServer(ctx context.Context, cfg *config.ApplicationConfig) {
	engine := gin.Default()
	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	engine.GET(IndexPath, func(c *gin.Context) {
		file, err := ioutil.ReadFile("C:\\数据\\codespace\\testgo\\cmd\\myshell\\index.html")
		if err != nil {
			panic(err.Error())
		}
		fmt.Fprint(c.Writer, string(file))
	})

	engine.GET("/ws", gin.WrapF(func(writer http.ResponseWriter, request *http.Request) {
		conn, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			panic(err.Error())
		}
		//TODO:id
		socketConn := NewWebSocketConn(ctx, conn, "test")
		go socketConn.Start()
	}))
	err := engine.Run(fmt.Sprintf(":%v", cfg.Port))
	if err != nil {
		panic(err.Error())
	}
}
