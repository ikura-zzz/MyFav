package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ユーザー登録画面への遷移
func RegistUser(engine *gin.Engine) {
	engine.GET("/registUser", func(c *gin.Context) {
		c.HTML(http.StatusOK, "registUser.html", gin.H{})
	})
}
