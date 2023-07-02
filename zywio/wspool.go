package zywio

import (
	"sync"
	"time"
)


type wsPool struct {
	// 心跳周期
	pingCycle time.Duration
	// 连接池
	pool sync.Map
}

func newWsPool(pingCycle time.Duration) *wsPool {
	wsp := &wsPool{
		pingCycle: pingCycle,
		pool: sync.Map{},
	}
	return wsp
}

// 往连接池加入一条新的连接
func (wp *wsPool) add(wsid string, ws *WsIO) {
	wp.pool.Store(wsid, ws)
}

// 从连接池删除一条连接
func (wp *wsPool) del(wsid string) {
	wp.pool.Delete(wsid)
}

// 向某连接发送数据
func (wp *wsPool) SendMessage(wsid string, data []byte) {

}

