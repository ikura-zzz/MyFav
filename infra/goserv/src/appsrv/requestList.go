package main

import (
	"fmt"
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
		transPage(c, switchlistalready)
	})
}
func Listnow(engine *gin.Engine) {
	engine.GET("/listnow", func(c *gin.Context) {
		transPage(c, switchlistnow)
	})
}
func Listwish(engine *gin.Engine) {
	engine.GET("/listwish", func(c *gin.Context) {
		transPage(c, switchlistwish)
	})
}

func switchlistall(c *gin.Context) {
	favs, err := genlistall(c)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.HTML(http.StatusOK, "listall.html", gin.H{
		"list": template.HTML(favs),
	})
}
func switchlistalready(c *gin.Context) {
	favs, err := genlistalready(c)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.HTML(http.StatusOK, "listalready.html", gin.H{
		"list": template.HTML(favs),
	})
}
func switchlistnow(c *gin.Context) {
	favs, err := genlistnow(c)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.HTML(http.StatusOK, "listnow.html", gin.H{
		"list": template.HTML(favs),
	})
}

func switchlistwish(c *gin.Context) {
	favs, err := genlistwish(c)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.HTML(http.StatusOK, "listwish.html", gin.H{
		"list": template.HTML(favs),
	})
}
