package main

import (
	"fmt"
	"html/template"
	"myfav/sessionmanager"
	"myfav/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(engine *gin.Engine) {
	engine.GET("/list", func(c *gin.Context) {
		username := c.Query("name")
		if username == "" {
			transPage(c, switchlist)
		} else {
			if sessionmanager.InqSessionValid(c, utils.SessionKeyUser) {
				switchlist(c)
			} else {
				switchlist_foreign(c, username)
			}
		}
	})
}

// 自分のリストへアクセスする場合
func switchlist(c *gin.Context) {
	favs, err := genlist(c)
	if err != nil {
		fmt.Println(err.Error())
	}
	username, _ := sessionmanager.GetSessionValue(c, utils.SessionKeyUser)
	c.HTML(http.StatusOK, "list.html", gin.H{
		"username": template.HTML(username),
		"list":     template.HTML(favs),
	})
}

// 他人のリストへアクセスする場合
func switchlist_foreign(c *gin.Context, username string) {
	favs, err := genlist_foreign(c, username)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.HTML(http.StatusOK, "list_foreign.html", gin.H{
		"username": template.HTML(username),
		"list":     template.HTML(favs),
	})
}
