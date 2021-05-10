package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web-tutorial/handler"
	"log"
)

func PrintGGG(context *gin.Context) {
	log.Printf("GGG")
}

func PrintGeo(context *gin.Context) {
	log.Printf("geo")
}

func PrintEarly(context *gin.Context) {
	log.Printf("early")
}

func main() {
	fmt.Println("Starting the server ...")
	defaultHandler := handler.NewDefaultHandler()
	helloHandler := handler.NewHelloHandler()
	// gin.Default()也适用gin.New()创建engine实例，但是会默认使用Logger和Recover中间件。
	engine := gin.Default()
	//在根节点使用GGG中间件
	engine.Use(PrintGGG)
	// 直接添加，该组使用PrintGeo中间件
	g := engine.Group("/g", PrintGeo)
	// 使用.GET等方法添加，会计算绝对请求路径，并添加给gin引擎
	// "/g"不会使用到PrintEarly中间件
	g.GET("/g", helloHandler.ServeHTTP)
	g.Use(PrintEarly)
	g.GET("/a", helloHandler.ServeHTTP)
	// engine也是一个Group
	engine.GET("/", defaultHandler.ServeHTTP)
	engine.GET("/hello", helloHandler.ServeHTTP)
	engine.GET("/a", func(context *gin.Context) {
		log.Printf("a接口")
	})
	// 默认运行在8080
	engine.Run()
	// 指定位置
	//engine.Run(":8081")
}
