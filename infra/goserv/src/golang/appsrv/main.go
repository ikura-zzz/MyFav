package main

import (
	"myfav/domain/logger"
	"myfav/domain/session"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	var l logger.Logger = new(logger.Logimp)
	l.Outlog("myfav started!!")
	engine := gin.Default()
	session.CreateNewSessionStore(engine)

	engine.Static("/static", "./static")
	engine.LoadHTMLGlob("/var/goserv/bin/html/*.html")
	setPages(engine)

	engine.Run(":8080")
}

// 各リクエストのハンドラを一斉起動
func setPages(engine *gin.Engine) {
	setAuth(engine)
	setFav(engine)
	setList(engine)
	setConfig(engine)
}

// GETリクエストのパラメータを全取得して返す。
func GetInputGET(c *gin.Context) []string {
	keys := strings.Split(c.Query("names"), ",")
	var inputs []string
	for i := 0; i < len(keys); i++ {
		inputs = append(inputs, c.Query(keys[i]))
	}
	return inputs
}

// POSTリクエストのパラメータを全取得して返す。
func GetInputPOST(c *gin.Context) []string {
	keys := strings.Split(c.PostForm("names"), ",")
	var inputs []string
	for i := 0; i < len(keys); i++ {
		inputs = append(inputs, c.PostForm(keys[i]))
	}
	return inputs
}
