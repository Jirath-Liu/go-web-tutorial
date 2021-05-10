package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Data struct {
	name string `json:"name"`
	age  int    `json:"age"`
}
type DefaultHandler struct{}

func NewDefaultHandler() *DefaultHandler {
	return &DefaultHandler{}
}

func (DefaultHandler) ServeHTTP(c *gin.Context) {
	data := Data{}
	c.ShouldBind(&data)
	c.JSON(http.StatusOK, data)
}
