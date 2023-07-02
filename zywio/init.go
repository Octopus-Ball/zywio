package zywio

import (
	// "encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"
	"github.com/gorilla/websocket"
)

type Req []byte
type Rsp []byte

type Context struct {
	wsID int
	wsio *WsIO
	req Req
	esp Rsp
}

func (c *Context) ReadJSON(obj any) error {

}

// ====================================================

type ReqHandler func(c *Context)

var idP int = 0
var pool *wsPool = nil
var reqHandler *ReqHandler

func getID() int {
	idP++
	return idP
}

// 设置websocket
// CheckOrigin防止跨站点的请求伪造
var upGrader = websocket.Upgrader {
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func InitWs(h ReqHandler) {
	pool = newWsPool()
	reqHandler = &h
}

func WsHandler(c *gin.Context) {
	id := getID()
	wsio, err := newWsIO(c)
	if err != nil {
		fmt.Println(err)
	}
	wsio.id = id
	pool.add(wsio.id, wsio)
}