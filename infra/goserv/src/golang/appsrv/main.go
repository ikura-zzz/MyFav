package main

import (
	"myfav/model/logger"
	"myfav/model/session"

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
