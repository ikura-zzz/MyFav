package main

import (
	"myfav/sessionmanager"
	"myfav/utils"

	"github.com/gin-gonic/gin"
)

func redirectTop(c *gin.Context) {
	c.Redirect(303, "/")
}

func redirectHome(c *gin.Context) {
	c.Redirect(303, "/list")
}

func transPage(c *gin.Context, fn func(*gin.Context)) {
	if !sessionmanager.InqSessionValid(c, utils.SessionKeyUser) {
		redirectTop(c)
		c.Abort()
	} else {
		fn(c)
	}
}