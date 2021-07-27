package main

import (
	"myfav/identifychk"
	"myfav/sessionmanager"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Config(engine *gin.Engine) {
	engine.POST("/config-username", func(c *gin.Context) {
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
