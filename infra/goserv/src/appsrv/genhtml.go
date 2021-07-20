package main

import (
	"errors"
	"myfav/favmanager"
	"myfav/sessionmanager"
	"myfav/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

const checked string = "checked"

func genlistall(c *gin.Context) (string, error) {
	return genlist(c, utils.SelectFavsByUserid)
}

func genlistalready(c *gin.Context) (string, error) {
	return genlist(c, utils.SelectFavsByUseridAlready)
}
func genlistnow(c *gin.Context) (string, error) {
	return genlist(c, utils.SelectFavsByUseridNow)
}

func genlistwish(c *gin.Context) (string, error) {
	return genlist(c, utils.SelectFavsByUseridWish)
}
func genlist(c *gin.Context, sql string) (string, error) {
	userid, err := sessionmanager.GetUserId(c)
	if err != nil {
		return "", errors.New("selectfav_getuserid " + err.Error())
	}
	favs, err := favmanager.SelectfavsByUserid(userid, sql)
	if err != nil {
		return "", errors.New("genlistall_selectfavs:" + err.Error())
	}
	return genhtml(favs)
}

func genhtml(favs []favmanager.Fav) (string, error) {
	html := ""
	for i := 0; i < len(favs); i++ {
		fav := favs[i]
		link := "/fav?favid=" + favs[i].Favid

		fav.Stars = starsconv(fav.Stars)

		html += "<a href=\"" + link +
			"\">	<li class=\"inline-block border-b border-gray-300 flex justify-between items-center py-4\">		<div class=\"flex items-start w-2/5\">			<div class=\"w-10 h-10 rounded mr-3\">				<img src=\"" +
			fav.Icon + "\" class=\"rounded-full h-10 w-10 bg-gray-300 m-auto\">			</div>			<div class=\"flex-1 overflow-hidden\">				<div>					<span class=\"font-bold \">" +
			fav.Title + "</span>				</div>				<p class=\"text-gray-700 text-xs\">" +
			fav.Category + "</p>			</div>		</div>		<p class=\"w-2/5\">" +
			fav.Overview + "</p>		<label for=\"stars\" class=\"font-bold w-1/5 text-right\">" +
			fav.Stars + "</label>	</li></a>\n"
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
		already + "><label class=\"tab_item text-base tracking-wide px-6\" for=\"already\">前から好き</label><input type=\"radio\" id=\"now\" name=\"timing\" value=\"now\" " +
		now + "><label class=\"tab_item text-base tracking-wide px-6\" for=\"now\">いま好き</label><input type=\"radio\" id=\"wish\" name=\"timing\" value=\"wish\" " +
		wish + "><label class=\"tab_item text-base tracking-wide px-6\" for=\"wish\">興味ある</label>"
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
		one + "><label class=\"tab_item text-base tracking-wide px-6\" for=\"one\">1</label><input type=\"radio\" id=\"two\" name=\"stars\" value=\"2\" " +
		two + "><label class=\"tab_item text-base tracking-wide px-6\" for=\"two\">2</label><input type=\"radio\" id=\"three\" name=\"stars\" value=\"3\" " +
		three + "><label class=\"tab_item text-base tracking-wide px-6\" for=\"three\">3</label><input type=\"radio\" id=\"four\" name=\"stars\" value=\"4\" " +
		four + "><label class=\"tab_item text-base tracking-wide px-6\" for=\"four\">4</label><input type=\"radio\" id=\"five\" name=\"stars\" value=\"5\" " +
		five + "><label class=\"tab_item text-base tracking-wide px-6\" for=\"five\">5</label>"
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
		open + "><label class=\"tab_item text-base tracking-wide px-6\" for=\"open\">公開</label><input type=\"radio\" id=\"close\" name=\"openclose\" value=\"close\" " +
		close + "><label class=\"tab_item text-base tracking-wide px-6\" for=\"close\">非公開</label>"
	return html
}
