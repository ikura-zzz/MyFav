package main

import (
	"errors"
	"fmt"
	"myfav/dbaccessor"
	"myfav/identifychk"
	"myfav/logmanager"
	"myfav/types"
	"myfav/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func genlist_foreign(c *gin.Context, username string) (string, error) {
	logmanager.Outlog("1")
	return genlistbody_foreign(utils.SelectFavsByUseridAndOpen, username)
}

func genlistbody_foreign(sql string, username string) (string, error) {
	logmanager.Outlog("2")
	userid, err := identifychk.GetUserId(username)
	if err != nil {
		return "", errors.New("genlistbody_foreign:" + err.Error())
	}
	logmanager.Outlog("3")
	favs, err := dbaccessor.Selectfavs(userid, sql)
	if err != nil {
		return "", errors.New("genlistbody_foreign_selectfav:" + err.Error())
	}
	return genhtml_foreign(favs, username)
}

func genhtml_foreign(favs []types.Fav, username string) (string, error) {
	html := ""
	for i := 0; i < len(favs); i++ {
		fav := favs[i]
		link := "/fav?favid=" + favs[i].Favid + "&name=" + username

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
	logmanager.Outlog("html" + html)
	return html, nil
}

func gentimingHTML_foreign(t string) string {
	var html string
	if t == "1" {
		html = "<input type=\"radio\" id=\"one\" name=\"stars\" value=\"前から\" checked><label class=\"tab_item text-base tracking-wide px-5\" for=\"one\">前から</label>"
	} else if t == "2" {
		html = "<input type=\"radio\" id=\"one\" name=\"stars\" value=\"今好き\" checked><label class=\"tab_item text-base tracking-wide px-5\" for=\"one\">今好き</label>"
	} else if t == "3" {
		html = "<input type=\"radio\" id=\"one\" name=\"stars\" value=\"興味ある\" checked><label class=\"tab_item text-base tracking-wide px-5\" for=\"one\">興味ある</label>"
	}

	return html
}

func genstarsHTML_foreign(stars string) string {
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
