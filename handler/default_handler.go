package handler

import (
	"bytes"
	"log"
	"net/http"
)

type DefaultHandler struct{}

func NewDefaultHandler() *DefaultHandler {
	return &DefaultHandler{}
}

func (DefaultHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	buf := bytes.Buffer{}
	buf.ReadFrom(request.Body)
	s := buf.String()
	log.Printf("\ndefault received \nMethod:%s\n%s\n%s\n", request.Method, request.RequestURI, s)
	writer.Write([]byte("hello"))
}
