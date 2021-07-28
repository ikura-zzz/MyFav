package main

import (
	"html/template"
	"myfav/identifychk"
	"myfav/logmanager"
	"myfav/sessionmanager"
	"myfav/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfigUser(engine *gin.Engine) {
	engine.GET("/configuser", func(c *gin.Context) {
		transPage(c, func(c *gin.Context) {
			c.HTML(http.StatusOK, "reconfPassword.html", gin.H{
				"action": template.HTML("/chgnameform"),
			})
		})
	})
}
func ConfigName(engine *gin.Engine) {
	engine.GET("/configpass", func(c *gin.Context) {
		transPage(c, func(c *gin.Context) {
			c.HTML(http.StatusOK, "reconfPassword.html", gin.H{
				"action": template.HTML("/chgpassform"),
			})
		})
	})
}
func ConfigDelUser(engine *gin.Engine) {
	engine.GET("/configuserdel", func(c *gin.Context) {
		transPage(c, func(c *gin.Context) {
			c.HTML(http.StatusOK, "reconfPassword.html", gin.H{
				"action": template.HTML("/userdeleteform"),
			})
		})
	})
}
func ConfigValidChk_name(engine *gin.Engine) {
	engine.POST("chgnameform", func(c *gin.Context) {
		username, ok := sessionmanager.GetSessionValue(c, utils.SessionKeyUser)
		if !ok {
			logmanager.Outlog("Config:can't get username:in configuser")
		}
		inputpass := c.PostForm("password")
		if err := identifychk.Invalidchk(username, inputpass); err != nil {
			transPage(c, func(c *gin.Context) {
				c.HTML(http.StatusOK, "reconfPassword.html", gin.H{
					"action": template.HTML("/chgnameform"),
					"errmsg": template.HTML(err.Error()),
				})
			})
			return
		}
		transPage(c, switchlist)
	})
}
func ConfigValidChk_pass(engine *gin.Engine) {
	engine.POST("chgpassform", func(c *gin.Context) {
		username, ok := sessionmanager.GetSessionValue(c, utils.SessionKeyUser)
		if !ok {
			logmanager.Outlog("Config:can't get username:in configpass")
		}
		inputpass := c.PostForm("password")
		if err := identifychk.Invalidchk(username, inputpass); err != nil {
			transPage(c, func(c *gin.Context) {
				c.HTML(http.StatusOK, "reconfPassword.html", gin.H{
					"action": template.HTML("/chgpassform"),
					"errmsg": template.HTML(err.Error()),
				})
			})
			return
		}
		transPage(c, switchlist)
	})
}
func ConfigValidChk_deluser(engine *gin.Engine) {
	engine.POST("userdeleteform", func(c *gin.Context) {
		username, ok := sessionmanager.GetSessionValue(c, utils.SessionKeyUser)
		if !ok {
			logmanager.Outlog("Config:can't get username:in configuserdel")
		}
		inputpass := c.PostForm("password")
		if err := identifychk.Invalidchk(username, inputpass); err != nil {
			transPage(c, func(c *gin.Context) {
				c.HTML(http.StatusOK, "reconfPassword.html", gin.H{
					"action": template.HTML("/userdeleteform"),
					"errmsg": template.HTML(err.Error()),
				})
			})
			return
		}
		transPage(c, switchlist)
	})
}
