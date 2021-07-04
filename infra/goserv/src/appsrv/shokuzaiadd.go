package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 食材登録画面への遷移
func Shokuzaiadd(engine *gin.Engine) {
	engine.GET("/shokuzaiadd", func(c *gin.Context) {
		c.HTML(http.StatusOK, "registShokuzai.html", gin.H{})
	})
}
