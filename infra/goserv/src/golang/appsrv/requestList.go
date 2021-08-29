package main

import (
	"html/template"
	"myfav/model/logger"
	"myfav/model/session"
	"myfav/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setList(engine *gin.Engine) {
	list(engine)
}

func list(engine *gin.Engine) {
	engine.GET("/list", func(c *gin.Context) {
		username := c.Query("name")
		if username == "" {
			transPage(c, switchlist)
		} else {
			switchlist_foreign(c, username)
		}
	})
}

// 自分のリストへアクセスする場合
func switchlist(c *gin.Context) {
	var l logger.Logger = new(logger.Logimp)
	favs, err := genlist(c)
	if err != nil {
		l.Outlog(err.Error())
	}
	username, _ := session.GetSessionValue(c, utils.SessionKeyUser)
	c.HTML(http.StatusOK, "list.html", gin.H{
		"username": template.HTML(username),
		"list":     template.HTML(favs),
	})
}

// 他人のリストへアクセスする場合
func switchlist_foreign(c *gin.Context, username string) {
	var l logger.Logger = new(logger.Logimp)
	favs, err := genlist_foreign(c, username)
	if err != nil {
		l.Outlog(err.Error())
	}
	c.HTML(http.StatusOK, "list_foreign.html", gin.H{
		"username": template.HTML(username),
		"list":     template.HTML(favs),
	})
}
