package main

import (
	"myfav/sessionmanager"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	sessionmanager.CreateNewSessionStore(engine)

	engine.Static("/static", "./static")
	engine.LoadHTMLGlob("/var/goserv/bin/html/*.html")
	setPages(engine)

	engine.Run(":8080")
}

// 各リクエストのハンドラを一斉起動
func setPages(engine *gin.Engine) {
	Signin(engine)
	Fav(engine)
	List(engine)
	Crtfav(engine)
	Modfav(engine)
	Delfav(engine)
	Signout(engine)
	Signup(engine)
	Config(engine)
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
