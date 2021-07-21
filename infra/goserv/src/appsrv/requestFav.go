package main

import (
	"fmt"
	"html/template"
	"myfav/favmanager"
	"myfav/logging"
	"myfav/sessionmanager"
	"myfav/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Crtfav(engine *gin.Engine) {
	engine.POST("/crtfav", func(c *gin.Context) {
		transPage(c, func(c *gin.Context) {
			fav := getPostElem(c)
			if err := favmanager.Favadd(c, fav); err != nil {
				logging.Log(err.Error(), logging.High)
			}
			redirectHome(c)
		})
	})
}

func Modfav(engine *gin.Engine) {
	engine.POST("/modfav", func(c *gin.Context) {
		transPage(c, func(c *gin.Context) {
			fav := getPostElem(c)
			if err := favmanager.Favmod(c, fav); err != nil {
				logging.Log(err.Error(), logging.High)
			}
			redirectHome(c)
		})
	})
}

func Delfav(engine *gin.Engine) {
	engine.POST("/delfav", func(c *gin.Context) {
		transPage(c, func(c *gin.Context) {
			fav := getPostElem(c)
			fmt.Println("favid:" + fav.Favid)
			if err := favmanager.Favdel(c, fav); err != nil {
				//logging.Log(err.Error(), logging.High)
				fmt.Println((err.Error()))
			}
			redirectHome(c)
		})
	})
}
func Fav(engine *gin.Engine) {
	engine.GET("/fav", func(c *gin.Context) {
		favid := c.Query("favid")
		userid, err := sessionmanager.GetUserId(c)
		if err != nil {
			transPage(c, func(c *gin.Context) {
				c.HTML(http.StatusOK, "fav.html", gin.H{})
			})
			return
		}
		favs, err := favmanager.Selectfavs(userid, utils.SelectFavsByUserid)
		if err != nil {
			logging.Log(err.Error(), logging.High)
			transPage(c, func(c *gin.Context) {
				c.HTML(http.StatusOK, "fav.html", gin.H{})
			})
			return
		}
		for _, f := range favs {
			if f.Favid == favid {
				transPage(c, func(c *gin.Context) {
					c.HTML(http.StatusOK, "fav_mod.html", gin.H{
						"icon":       template.URL(f.Icon),
						"title":      template.HTML(f.Title),
						"category":   template.HTML(f.Category),
						"publisher":  template.HTML(f.Publisher),
						"overview":   template.HTML(f.Overview),
						"impression": template.HTML(f.Impre),
						"timing":     template.HTML(gentimingHTML(f.Timing)),
						"stars":      template.HTML(genstarsHTML(f.Stars)),
						"openclose":  template.HTML(genopencloseHTML(f.Openclose)),
						"favid":      template.HTML(f.Favid),
					})
				})
				return
			}
		}
		transPage(c, func(c *gin.Context) {
			c.HTML(http.StatusOK, "fav.html", gin.H{})
		})
	})
}

func getPostElem(c *gin.Context) favmanager.Fav {
	var fav favmanager.Fav
	fav.Favid = c.PostForm("favid")
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
