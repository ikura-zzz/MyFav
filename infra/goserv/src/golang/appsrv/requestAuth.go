package main

import (
	"myfav/domain/session"
	"myfav/domain/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setAuth(engine *gin.Engine) {
	Signin(engine)
	Signout(engine)
	Signup(engine)
}

func Signin(engine *gin.Engine) {
	engine.POST("/signin", func(c *gin.Context) {
		errtrans := func(msg string) {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"errmsg": msg,
			})
		}
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username == "" || password == "" {
			errtrans("ユーザー名またはパスワードが未入力です。")
			return
		}
		if err := user.Invalidchk(username, password); err != nil {
			errtrans(err.Error())
			return
		} else {
			session.CreateNewSession(c, username)
			redirectHome(c)
		}
	})
}
func Signout(engine *gin.Engine) {
	engine.GET("/signout", func(c *gin.Context) {
		session.RemoveSession(c)
		redirectTop(c)
	})
}
func Signup(engine *gin.Engine) {
	engine.POST("/signup", func(c *gin.Context) {
		errtrans := func(msg string) {
			c.HTML(http.StatusOK, "newuser.html", gin.H{
				"errmsg": msg,
			})
		}
		username := c.PostForm("username")
		password := c.PostForm("password")
		retype := c.PostForm("retypepassword")
		if username == "" || password == "" || retype == "" {
			errtrans("ユーザー名またはパスワードが未入力です。")
			return
		}
		if err := user.Useradd(username, password, retype); err != nil {
			errtrans(err.Error())
			return
		} else {
			redirectTop(c)
		}
	})
}
