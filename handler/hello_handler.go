package handler

import (
	"github.com/gin-gonic/gin"
)

type HelloHandler struct {
	m map[string]string
}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{m: make(map[string]string)}
}

// 回声服务器，返回接受的body
// 实现Handler接口
func (h HelloHandler) ServeHTTP(c *gin.Context) {
	//c.ShouldBindUri()
}
