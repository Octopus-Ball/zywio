package main

import (
	"fmt"
	"sync"
	// "zy/zywio"

	"github.com/gin-gonic/gin"
)

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

func pingHandler(c *gin.Context) {
	c.String(200, "pong")
}


func main() {
	// g := gin.Default()
	// wio := zywio.Default()

	// g.GET("/ping", pingHandler)
	// g.GET("/ws", zywio.WsHandler)

	// g.Run(":8888")

	var m sync.Map

	v, t := m.LoadOrStore("k1", "v1")
	fmt.Println(v)
	fmt.Println(t)

	v1, t1 := m.LoadOrStore("k1", "v2")
	fmt.Println(v1)
	fmt.Println(t1)
}