package main

import (
	"html/template"
	"myfav/favmanager"
	"myfav/identifychk"
	"myfav/logmanager"
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
				logmanager.Outlog(err.Error())
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
				logmanager.Outlog(err.Error())
			}
			redirectHome(c)
		})
	})
}

func Delfav(engine *gin.Engine) {
	engine.POST("/delfav", func(c *gin.Context) {
		transPage(c, func(c *gin.Context) {
			fav := getPostElem(c)
			logmanager.Outlog("favid:" + fav.Favid)
			if err := favmanager.Favdel(c, fav); err != nil {
				logmanager.Outlog((err.Error()))
			}
			redirectHome(c)
		})
	})
}
func Fav(engine *gin.Engine) {
	engine.GET("/fav", func(c *gin.Context) {
		favid := c.Query("favid")
		username := c.Query("name")
		var userid int
		var err error
		if username != "" {
			userid, err = identifychk.GetUserId(username)
			if err != nil {
				transPage(c, func(c *gin.Context) {
					c.HTML(http.StatusOK, "fav.html", gin.H{})
				})
				return
			}
		} else {
			userid, err = sessionmanager.GetUserId(c)
			if err != nil {
				transPage(c, func(c *gin.Context) {
					c.HTML(http.StatusOK, "fav.html", gin.H{})
				})
				return
			}
		}
		favs, err := favmanager.Selectfavs(userid, utils.SelectFavsByUserid)
		if err != nil {
			logmanager.Outlog(err.Error())
			transPage(c, func(c *gin.Context) {
				c.HTML(http.StatusOK, "fav.html", gin.H{})
			})
			return
		}

		if username != "" {
			favforeigner(c, favs, favid, username)
			return
		} else if favid == "" {
			transPage(c, func(c *gin.Context) {
				c.HTML(http.StatusOK, "fav.html", gin.H{})
			})
		} else {
			myfavsend(c, favs, favid)
			return
		}
	})
}
func favforeigner(c *gin.Context, favs []favmanager.Fav, favid string, username string) {

	for _, f := range favs {
		if f.Favid == favid {
			c.HTML(http.StatusOK, "fav_foreign.html", gin.H{
				"username":   template.HTML(username),
				"icon":       template.URL(f.Icon),
				"title":      template.HTML(f.Title),
				"category":   template.HTML(f.Category),
				"publisher":  template.HTML(f.Publisher),
				"overview":   template.HTML(f.Overview),
				"impression": template.HTML(f.Impre),
				"timing":     template.HTML(gentimingHTML_foreign(f.Timing)),
				"stars":      template.HTML(genstarsHTML_foreign(f.Stars)),
				"favid":      template.HTML(f.Favid),
			})
			return
		}
	}
}
func myfavsend(c *gin.Context, favs []favmanager.Fav, favid string) {

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
