package handler

import (
	"fmt"
	"net/http"
)

type HelloHandler struct {
	m map[string]string
}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{m: make(map[string]string)}
}

// 回声服务器，返回接受的body
func (h HelloHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	b := request.Body
	buf := make([]byte, 5000)
	len, _ := b.Read(buf)
	b.Close()
	s := string(buf)
	fmt.Printf("get request: \n%s\n%s\n", request.RequestURI, s)
	writer.Write(buf[:len])
}
