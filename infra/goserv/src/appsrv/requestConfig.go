package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Config(engine *gin.Engine) {
	engine.POST("/config-username", func(c *gin.Context) {
		transPage(c, func(c *gin.Context) {
			c.HTML(http.StatusOK, "reconfPassword.html", gin.H{
				"action": template.HTML("chgnameform"),
			})
		})
	})
	engine.POST("/config-password", func(c *gin.Context) {
		transPage(c, func(c *gin.Context) {
			c.HTML(http.StatusOK, "reconfPassword.html", gin.H{
				"action": template.HTML("chgpassform"),
			})
		})
	})
	engine.POST("/config-userdelete", func(c *gin.Context) {
		transPage(c, func(c *gin.Context) {
			c.HTML(http.StatusOK, "reconfPassword.html", gin.H{
				"action": template.HTML("userdeleteform"),
			})
		})
	})
}
