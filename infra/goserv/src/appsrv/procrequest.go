package main

import (
	"fmt"
	"html/template"
	"myfav/crtusr"
	"myfav/identifychk"
	"myfav/sessionmanager"
	"myfav/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ログイン要求
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
			switchlistall(c)
		}
	})
}
func Signout(engine *gin.Engine) {
	engine.GET("/signout", func(c *gin.Context) {
		sessionmanager.RemoveSession(c)
		redirectHome(c)
	})
}
func Signup(engine *gin.Engine) {
	engine.POST("/signup", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		retype := c.PostForm("retypepassword")
		fmt.Println(username)
		fmt.Println(password)
		fmt.Println(retype)
		if err := crtusr.Useradd(username, password, retype); err != nil {
			c.HTML(http.StatusOK, "newuser.html", gin.H{
				"errmsg": err.Error(),
			})
			return
		} else {
			redirectHome(c)
		}
	})
}
func Listall(engine *gin.Engine) {
	engine.GET("/listall", func(c *gin.Context) {
		transPage(c, switchlistall)
	})
}
func Listalready(engine *gin.Engine) {
	engine.GET("/listalready", func(c *gin.Context) {
		transPage(c, func(c *gin.Context) {
			c.HTML(http.StatusOK, "listalready.html", gin.H{})
		})
	})
}
func Listnow(engine *gin.Engine) {
	engine.GET("/listnow", func(c *gin.Context) {
		transPage(c, func(c *gin.Context) {
			c.HTML(http.StatusOK, "listnow.html", gin.H{})
		})
	})
}
func Listwish(engine *gin.Engine) {
	engine.GET("/listwish", func(c *gin.Context) {
		transPage(c, func(c *gin.Context) {
			c.HTML(http.StatusOK, "listwish.html", gin.H{})
		})
	})
}
func Crtfav(engine *gin.Engine) {
	engine.POST("/crtfav", func(c *gin.Context) {
		transPage(c, func(c *gin.Context) {
			switchlistall(c)
		})
	})
}
func Fav(engine *gin.Engine) {
	engine.GET("/fav", func(c *gin.Context) {
		transPage(c, func(c *gin.Context) {
			c.HTML(http.StatusOK, "fav.html", gin.H{})
		})
	})
}
func switchlistall(c *gin.Context) {
	c.HTML(http.StatusOK, "listall.html", gin.H{
		"list": template.HTML(genlistall()),
	})
}

func redirectHome(c *gin.Context) {
	c.Redirect(303, "/")
}

func transPage(c *gin.Context, fn func(*gin.Context)) {
	if !sessionmanager.InqSessionValid(c, utils.SessionKeyUser) {
		redirectHome(c)
		c.Abort()
	} else {
		fn(c)
	}
}
