package main

import (
	"html/template"
	"myfav/logging"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Listall(engine *gin.Engine) {
	engine.GET("/listall", func(c *gin.Context) {
		username := c.Query("name")
		if username == "" {
			transPage(c, switchlistall)
		} else {
			switchlistall_foreign(c)
		}
	})
}
func Listalready(engine *gin.Engine) {
	engine.GET("/listalready", func(c *gin.Context) {
		username := c.Query("name")
		if username == "" {
			transPage(c, switchlistalready)
		} else {
			switchlistalready_foreign(c)
		}
	})
}
func Listnow(engine *gin.Engine) {
	engine.GET("/listnow", func(c *gin.Context) {
		username := c.Query("name")
		if username == "" {
			transPage(c, switchlistnow)
		} else {
			switchlistnow_foreign(c)
		}
	})
}
func Listwish(engine *gin.Engine) {
	engine.GET("/listwish", func(c *gin.Context) {
		username := c.Query("name")
		if username == "" {
			transPage(c, switchlistwish)
		} else {
			switchlistwish_foreign(c)
		}
	})
}

// 自分のリストへアクセスする場合
func switchlistall(c *gin.Context) {
	favs, err := genlistall(c)
	if err != nil {
		logging.Log(err.Error(), logging.High)
	}
	c.HTML(http.StatusOK, "listall.html", gin.H{
		"list": template.HTML(favs),
	})
}
func switchlistalready(c *gin.Context) {
	favs, err := genlistalready(c)
	if err != nil {
		logging.Log(err.Error(), logging.High)
	}
	c.HTML(http.StatusOK, "listalready.html", gin.H{
		"list": template.HTML(favs),
	})
}
func switchlistnow(c *gin.Context) {
	favs, err := genlistnow(c)
	if err != nil {
		logging.Log(err.Error(), logging.High)
	}
	c.HTML(http.StatusOK, "listnow.html", gin.H{
		"list": template.HTML(favs),
	})
}
func switchlistwish(c *gin.Context) {
	favs, err := genlistwish(c)
	if err != nil {
		logging.Log(err.Error(), logging.High)
	}
	c.HTML(http.StatusOK, "listwish.html", gin.H{
		"list": template.HTML(favs),
	})
}

// 他人のリストへアクセスする場合
func switchlistall_foreign(c *gin.Context) {
	favs, err := genlistall_foreign(c)
	if err != nil {
		logging.Log(err.Error(), logging.High)
	}
	c.HTML(http.StatusOK, "listall.html", gin.H{
		"list": template.HTML(favs),
	})
}
func switchlistalready_foreign(c *gin.Context) {
	favs, err := genlistalready_foreign(c)
	if err != nil {
		logging.Log(err.Error(), logging.High)
	}
	c.HTML(http.StatusOK, "listalready.html", gin.H{
		"list": template.HTML(favs),
	})
}
func switchlistnow_foreign(c *gin.Context) {
	favs, err := genlistnow_foreign(c)
	if err != nil {
		logging.Log(err.Error(), logging.High)
	}
	c.HTML(http.StatusOK, "listnow.html", gin.H{
		"list": template.HTML(favs),
	})
}
func switchlistwish_foreign(c *gin.Context) {
	favs, err := genlistwish_foreign(c)
	if err != nil {
		logging.Log(err.Error(), logging.High)
	}
	c.HTML(http.StatusOK, "listwish.html", gin.H{
		"list": template.HTML(favs),
	})
}
