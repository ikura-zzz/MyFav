package main

import (
	"html/template"
	"myfav/dbaccessor"
	"myfav/model/fav"
	"myfav/model/logger"
	"myfav/model/session"
	"myfav/types"
	"myfav/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setFav(engine *gin.Engine) {
	favPage(engine)
	crtfav(engine)
	modfav(engine)
	delfav(engine)
}

func crtfav(engine *gin.Engine) {
	var l logger.Logger = new(logger.Logimp)
	engine.POST("/crtfav", func(c *gin.Context) {
		transPage(c, func(c *gin.Context) {
			f := getPostElem(c)
			var err error
			f.Userid, err = session.GetUserId(c)
			if err != nil {
				l.Outlog("favadd getuserid:" + err.Error())
			}
			if err := fav.Favadd(f); err != nil {
				l.Outlog(err.Error())
			}
			redirectHome(c)
		})
	})
}

func modfav(engine *gin.Engine) {
	var l logger.Logger = new(logger.Logimp)
	engine.POST("/modfav", func(c *gin.Context) {
		transPage(c, func(c *gin.Context) {
			f := getPostElem(c)
			var err error
			f.Userid, err = session.GetUserId(c)
			if err != nil {
				l.Outlog("favmod getuserid:" + err.Error())
			}
			if err := fav.Favmod(f); err != nil {
				l.Outlog(err.Error())
			}
			redirectHome(c)
		})
	})
}

func delfav(engine *gin.Engine) {
	var l logger.Logger = new(logger.Logimp)
	engine.POST("/delfav", func(c *gin.Context) {
		transPage(c, func(c *gin.Context) {
			f := getPostElem(c)
			l.Outlog("favid:" + f.Favid)
			var err error
			f.Userid, err = session.GetUserId(c)
			if err != nil {
				l.Outlog("favdel getuserid:" + err.Error())
			}
			if err := fav.Favdel(f); err != nil {
				l.Outlog((err.Error()))
			}
			redirectHome(c)
		})
	})
}
func favPage(engine *gin.Engine) {
	var l logger.Logger = new(logger.Logimp)
	engine.GET("/fav", func(c *gin.Context) {
		favid := c.Query("favid")
		username := c.Query("name")
		var userid int
		var err error
		if username != "" {
			userid, err = dbaccessor.GetUserId(username)
			if err != nil {
				transPage(c, func(c *gin.Context) {
					c.HTML(http.StatusOK, "fav.html", gin.H{})
				})
				return
			}
		} else {
			userid, err = session.GetUserId(c)
			if err != nil {
				transPage(c, func(c *gin.Context) {
					c.HTML(http.StatusOK, "fav.html", gin.H{})
				})
				return
			}
		}
		favs, err := dbaccessor.Selectfavs(userid, utils.SelectFavsByUserid)
		if err != nil {
			l.Outlog(err.Error())
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
func favforeigner(c *gin.Context, favs []types.Fav, favid string, username string) {

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
func myfavsend(c *gin.Context, favs []types.Fav, favid string) {

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

func getPostElem(c *gin.Context) types.Fav {
	var fav types.Fav
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
