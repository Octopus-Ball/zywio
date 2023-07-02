// package zywio

// import (
// 	"fmt"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/gorilla/websocket"
// )





// func NewZywio(reqHandler ReqHandler) *Zywio {
// 	return &Zywio{
// 		wsPool: make(map[int]*websocket.Conn),
// 		reqHandler: reqHandler,
// 	}
// }

// func (zwio *Zywio) read(handlerMsg ReqHandler, wsID int) {
// 	ws := zwio.wsPool[wsID]
// 	for {
// 		reqMsg := new(wsMsg)
// 		err := ws.ReadJSON(reqMsg)
// 		if err != nil {
// 			fmt.Println("read 错误")
// 			fmt.Println(err)
// 			break
// 		}
// 		rspMsg := handlerMsg(reqMsg)
// 		if rspMsg != nil {
// 			zwio.WriteByWS(ws, rspMsg)
// 		}
// 	}
// }

// func (zwio *Zywio) WriteByWS(ws *websocket.Conn, rspMsg *wsMsg) error {
// 	return ws.WriteJSON(rspMsg)
// }

// func (zwio *Zywio) WriteByWSID(wsID int, rspMsg *wsMsg) error {
// 	ws := zwio.wsPool[wsID]
// 	return zwio.WriteByWS(ws, rspMsg)
// }

// func (zwio *Zywio) WsHandler(c *gin.Context) {
// 	wsID := getID()
// 	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		ws.Close() //返回前关闭
// 		return
// 	}

// 	zwio.wsPool[wsID] = ws
// 	go func() {
// 		zwio.read(zwio.reqHandler, wsID)
// 	}()
// }