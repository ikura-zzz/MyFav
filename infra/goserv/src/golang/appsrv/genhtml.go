package main

import (
	"errors"
	"fmt"
	"myfav/dbaccessor"
	"myfav/domain/session"
	"myfav/types"
	"myfav/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

const checked string = "checked"

func genlist(c *gin.Context) (string, error) {
	userid, err := session.GetUserId(c)
	if err != nil {
		return "", errors.New("genlistbody:" + err.Error())
	}
	favs, err := dbaccessor.Selectfavs(userid, utils.SelectFavsByUserid)
	if err != nil {
		return "", errors.New("genlistbody_selectfav:" + err.Error())
	}
	return genhtml(favs)
}

func genhtml(favs []types.Fav) (string, error) {
	html := ""
	for i := 0; i < len(favs); i++ {
		fav := favs[i]
		link := "/fav?favid=" + favs[i].Favid
		if fav.Icon == "" {
			fav.Icon = "https://www.s-myfav.com/icons/apple-touch-icon.png"
		}

		fav.Stars = starsconv(fav.Stars)

		html += "<a id=\"fav" + fmt.Sprintf("%d", i) + "\" href=\"" + link + "\" style=\"display: none;\">" +
			"<input type=\"hidden\" id=\"timing" + fmt.Sprintf("%d", i) + "\" value=\"" + fav.Timing + "\">" +
			"<li class=\"inline-block border-b border-gray-300 flex justify-between items-center py-4\" >" +
			"<div class=\"w-10 h-10 rounded mr-3\">" +
			"<img src=\"" + fav.Icon + "\" class=\"rounded-full h-10 w-10 bg-gray-300 m-auto\">" +
			"</div>" +
			"<div class=\"flex-1 items-start\">" +
			"<span class=\"font-bold \">" + fav.Title + "</span>" +
			"<div class=\"flex overflow-hidden\">" +
			"<p class=\"flex-1 text-gray-700 text-xs\">" + fav.Category + "</p>" +
			"<p class=\"w-3/5\">" + fav.Overview + "</p>" +
			"<label for=\"stars\" class=\"text-right\">" + fav.Stars + "</label>" +
			"</div>" +
			"</div></li></a>\n"
	}
	return html, nil
}

func starsconv(stars string) string {
	starnum, _ := strconv.Atoi(stars)
	starstr := ""
	for i := 0; i < 5; i++ {
		if i < starnum {
			starstr += "★"
		} else {
			starstr += "☆"
		}
	}
	return starstr
}
func gentimingHTML(t string) string {
	var already, now, wish string
	already = ""
	now = ""
	wish = ""
	if t == "1" {
		already = checked
	} else if t == "2" {
		now = checked
	} else if t == "3" {
		wish = checked
	}
	html := "<input type=\"radio\" id=\"already\" name=\"timing\" value=\"already\" " +
		already + "><label class=\"tab_item text-base tracking-wide px-3\" for=\"already\">前から</label><input type=\"radio\" id=\"now\" name=\"timing\" value=\"now\" " +
		now + "><label class=\"tab_item text-base tracking-wide px-3\" for=\"now\">今好き</label><input type=\"radio\" id=\"wish\" name=\"timing\" value=\"wish\" " +
		wish + "><label class=\"tab_item text-base tracking-wide px-3\" for=\"wish\">興味ある</label>"
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
		one = checked
	} else if s == "2" {
		two = checked
	} else if s == "3" {
		three = checked
	} else if s == "4" {
		four = checked
	} else if s == "5" {
		five = checked
	}
	html := "<input type=\"radio\" id=\"one\" name=\"stars\" value=\"1\" " +
		one + "><label class=\"tab_item text-base tracking-wide px-5\" for=\"one\">1</label><input type=\"radio\" id=\"two\" name=\"stars\" value=\"2\" " +
		two + "><label class=\"tab_item text-base tracking-wide px-5\" for=\"two\">2</label><input type=\"radio\" id=\"three\" name=\"stars\" value=\"3\" " +
		three + "><label class=\"tab_item text-base tracking-wide px-5\" for=\"three\">3</label><input type=\"radio\" id=\"four\" name=\"stars\" value=\"4\" " +
		four + "><label class=\"tab_item text-base tracking-wide px-5\" for=\"four\">4</label><input type=\"radio\" id=\"five\" name=\"stars\" value=\"5\" " +
		five + "><label class=\"tab_item text-base tracking-wide px-5\" for=\"five\">5</label>"
	return html
}
func genopencloseHTML(oc string) string {
	var open, close string
	open = ""
	close = ""
	if oc == "1" {
		open = checked
	} else if oc == "0" {
		close = checked
	}
	html := "<input type=\"radio\" id=\"open\" name=\"openclose\" value=\"open\" " +
		open + "><label class=\"tab_item text-base tracking-wide px-10\" for=\"open\">公開</label><input type=\"radio\" id=\"close\" name=\"openclose\" value=\"close\" " +
		close + "><label class=\"tab_item text-base tracking-wide px-10\" for=\"close\">非公開</label>"
	return html
}
