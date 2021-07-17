package main

import (
	"fmt"
	"myfav/favmanager"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Crtfav(engine *gin.Engine) {
	engine.POST("/crtfav", func(c *gin.Context) {
		transPage(c, func(c *gin.Context) {
			fav := getPostElem(c)
			if err := favmanager.Favadd(c, fav); err != nil {
				fmt.Println(err)
			}
			redirectHome(c)
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

func getPostElem(c *gin.Context) favmanager.Fav {
	var fav favmanager.Fav
	fav.Icon = c.PostForm("icon")
	fav.Title = c.PostForm("title")
	fav.Category = c.PostForm("category")
	fav.Publisher = c.PostForm("publisher")
	fav.Overview = c.PostForm("overview")
	fav.Impre = c.PostForm("impre")
	fav.Timing = c.PostForm("timing")
	fav.Stars = c.PostForm("stars")
	fav.Openclose = c.PostForm("openclose")
	return fav
}
