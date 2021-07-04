package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// トップページ
func Index(engine *gin.Engine) {
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
}
