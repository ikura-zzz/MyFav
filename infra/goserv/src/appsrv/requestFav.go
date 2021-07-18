package main

import (
	"fmt"
	"html/template"
	"myfav/favmanager"
	"myfav/utils"
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

func Modfav(engine *gin.Engine) {
	engine.POST("/modfav", func(c *gin.Context) {
		transPage(c, func(c *gin.Context) {
			fav := getPostElem(c)
			if err := favmanager.Favmod(c, fav); err != nil {
				fmt.Println(err)
			}
			redirectHome(c)
		})
	})
}
func Fav(engine *gin.Engine) {
	engine.GET("/fav", func(c *gin.Context) {
		favid := c.Query("favid")
		userid, err := favmanager.Getusrid(c)
		if err != nil {
			fmt.Println("switchFav_getusrid:" + err.Error())
		}
		favs, err := favmanager.SelectfavsByUserid(userid, utils.SelectFavsByUserid)
		if err != nil {
			fmt.Println("switchFav:" + err.Error())
		}
		for _, f := range favs {
			if f.Favid == favid {
				transPage(c, func(c *gin.Context) {
					c.HTML(http.StatusOK, "fav_mod.html", gin.H{
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
				c.Abort()
			}
		}
		transPage(c, func(c *gin.Context) {
			c.HTML(http.StatusOK, "fav.html", gin.H{})
		})
	})
}

func gentimingHTML(t string) string {
	var already, now, wish string
	already = ""
	now = ""
	wish = ""
	if t == "1" {
		already = "checked"
	} else if t == "2" {
		now = "checked"
	} else if t == "3" {
		wish = "checked"
	}
	html := "<input type=\"radio\" id=\"already\" name=\"timing\" value=\"already\" " + already + "><label class=\"tab_item text-base tracking-wide px-6\" for=\"already\">前から好き</label><input type=\"radio\" id=\"now\" name=\"timing\" value=\"now\" " + now + "><label class=\"tab_item text-base tracking-wide px-6\" for=\"now\">いま好き</label><input type=\"radio\" id=\"wish\" name=\"timing\" value=\"wish\" " + wish + "><label class=\"tab_item text-base tracking-wide px-6\" for=\"wish\">好きになりたい</label>"
	return html
}

func genstarsHTML(s string) string {
	var one, two, three, four, five string
	one = ""
	two = ""
	three = ""
	four = ""
	three = ""
	if s == "1" {
		one = "checked"
	} else if s == "2" {
		two = "checked"
	} else if s == "3" {
		three = "checked"
	} else if s == "4" {
		four = "checked"
	} else if s == "5" {
		five = "chkecked"
	}
	html := "<input type=\"radio\" id=\"one\" name=\"stars\" value=\"1\" " + one + "><label class=\"tab_item text-base tracking-wide px-6\" for=\"one\">1</label><input type=\"radio\" id=\"two\" name=\"stars\" value=\"2\" " + two + "><label class=\"tab_item text-base tracking-wide px-6\" for=\"two\">2</label><input type=\"radio\" id=\"three\" name=\"stars\" value=\"3\" " + three + "><label class=\"tab_item text-base tracking-wide px-6\" for=\"three\">3</label><input type=\"radio\" id=\"four\" name=\"stars\" value=\"4\" " + four + "><label class=\"tab_item text-base tracking-wide px-6\" for=\"four\">4</label><input type=\"radio\" id=\"five\" name=\"stars\" value=\"5\" " + five + "><label class=\"tab_item text-base tracking-wide px-6\" for=\"five\">5</label>"
	return html
}
func genopencloseHTML(oc string) string {
	var open, close string
	open = ""
	close = ""
	if oc == "1" {
		open = "checked"
	} else if oc == "2" {
		close = "checked"
	}
	html := "<input type=\"radio\" id=\"open\" name=\"openclose\" value=\"open\" " + open + "><label class=\"tab_item text-base tracking-wide px-6\" for=\"open\">公開</label><input type=\"radio\" id=\"close\" name=\"openclose\" value=\"close\" " + close + "<label class=\"tab_item text-base tracking-wide px-6\" for=\"close\">非公開<label>"
	return html
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
