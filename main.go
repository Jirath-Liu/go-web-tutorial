package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-web-tutorial/handler"
	"log"
	"net/http"
)

func AuthCheck(handle httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		log.Printf("Auth check ")
		log.Printf(params.ByName("auth"))
		if params.ByName("auth") != "" {
			handle(writer, request, params)
		} else {
			http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}

	}
}

func main() {
	fmt.Println("Starting the server ...")

	defaultHandler := handler.NewDefaultHandler()
	helloHandler := handler.NewHelloHandler()

	router := httprouter.New()
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", r.Header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})
	router.GET("/", AuthCheck(defaultHandler.ServeHTTP))
	router.GET("/hello/:name", helloHandler.ServeHTTP)
	router.POST("/hello/", helloHandler.ServeHTTP)
	router.PUT("hello", helloHandler.ServeHTTP)
	router.PATCH("patch", helloHandler.ServeHTTP)
	router.DELETE("/hello", defaultHandler.ServeHTTP)

	log.Fatal(http.ListenAndServe(":8080", router))
}
