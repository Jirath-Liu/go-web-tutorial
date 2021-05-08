package handler

import (
	"bytes"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type DefaultHandler struct{}

func NewDefaultHandler() *DefaultHandler {
	return &DefaultHandler{}
}

func (DefaultHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	buf := bytes.Buffer{}
	buf.ReadFrom(request.Body)
	s := buf.String()
	log.Printf("\ndefault received \nMethod:%s\n%s\n%s\n", request.Method, request.RequestURI, s)
	writer.Write([]byte("hello"))
}
