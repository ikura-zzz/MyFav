package main

import (
	"myfav/crtusr"
	"myfav/identifychk"
	"myfav/sessionmanager"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signin(engine *gin.Engine) {
	engine.POST("/signin", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if err := identifychk.Invalidchk(username, password); err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"errmsg": err.Error(),
			})
		} else {
			sessionmanager.CreateNewSession(c, username)
			redirectHome(c)
		}
	})
}
func Signout(engine *gin.Engine) {
	engine.GET("/signout", func(c *gin.Context) {
		sessionmanager.RemoveSession(c)
		redirectTop(c)
	})
}
func Signup(engine *gin.Engine) {
	engine.POST("/signup", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		retype := c.PostForm("retypepassword")
		if err := crtusr.Useradd(username, password, retype); err != nil {
			c.HTML(http.StatusOK, "newuser.html", gin.H{
				"errmsg": err.Error(),
			})
			return
		} else {
			redirectTop(c)
		}
	})
}
