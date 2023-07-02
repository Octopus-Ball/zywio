package zywio

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WsIO struct {
	id int
	ws *websocket.Conn
}

func newWsIO(c *gin.Context) (wsio *WsIO, err error) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return nil, err
	}
	wsio = new(WsIO)
	wsio.ws = ws
	return
}

func (w *WsIO) read(v interface{}) error {
	return w.ws.ReadJSON(v)
}

func (w *WsIO) write(v interface{}) error {
	w.ws.WriteMessage()
	return w.ws.WriteJSON(v)
}

func (w *WsIO) close() {
	w.ws.Close()
}

func (w *WsIO) loopRead() {
	for {
		w.ws.ReadMessage()
	}
}