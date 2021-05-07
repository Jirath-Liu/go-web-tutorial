package main

import (
	"fmt"
	"go-web-tutorial/handler"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting the server ...")

	http.Handle("/", handler.NewDefaultHandler())
	helloHandler := handler.NewHelloHandler()
	http.Handle("/hello", helloHandler)
	// 创建服务器，ListenAndServe若服务器宕机，会返回异常
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
