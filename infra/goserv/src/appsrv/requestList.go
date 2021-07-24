package main

import (
	"html/template"
	"myfav/logging"
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(engine *gin.Engine) {
	engine.GET("/list", func(c *gin.Context) {
		username := c.Query("name")
		if username == "" {
			transPage(c, switchlist)
		} else {
			switchlist_foreign(c)
		}
	})
}

// 自分のリストへアクセスする場合
func switchlist(c *gin.Context) {
	favs, err := genlist(c)
	if err != nil {
		logging.Log(err.Error(), logging.High)
	}
	c.HTML(http.StatusOK, "list.html", gin.H{
		"list": template.HTML(favs),
	})
}

// 他人のリストへアクセスする場合
func switchlist_foreign(c *gin.Context) {
	favs, err := genlist_foreign(c)
	if err != nil {
		logging.Log(err.Error(), logging.High)
	}
	c.HTML(http.StatusOK, "list.html", gin.H{
		"list": template.HTML(favs),
	})
}
