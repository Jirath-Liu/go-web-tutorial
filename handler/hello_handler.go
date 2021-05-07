package handler

import (
	"bytes"
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
// 实现Handler接口
func (h HelloHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	b := request.Body
	buf := bytes.Buffer{}
	buf.ReadFrom(b)
	b.Close()
	s := buf.String()
	fmt.Printf("get request: \nMethod:%s\n%s\n%s\n", request.Method, request.RequestURI, s)
	writer.Write(buf.Bytes())
}
