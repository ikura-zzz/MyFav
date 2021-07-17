package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func switchlistall(c *gin.Context) {
	c.HTML(http.StatusOK, "listall.html", gin.H{
		"list": template.HTML(genlistall()),
	})
}
