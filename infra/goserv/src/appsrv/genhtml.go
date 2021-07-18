package main

import (
	"errors"
	"fmt"
	"myfav/favmanager"
	"myfav/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
	userid, err := favmanager.Getusrid(c)
	if err != nil {
		return "", errors.New("selectfav_getusrid " + err.Error())
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
		link := fmt.Sprintf("/fav?favid=%d", favs[i].Favid)

		fav.Stars = starsconv(fav.Stars)

		html += "<a href=\"" + link + "\">	<li class=\"inline-block border-b border-gray-300 flex justify-between items-center py-4\">		<div class=\"flex items-start w-2/5\">			<div class=\"w-10 h-10 rounded mr-3\">				<img src=\"" + fav.Icon + "\" class=\"rounded-full h-10 w-10 bg-gray-300 m-auto\">			</div>			<div class=\"flex-1 overflow-hidden\">				<div>					<span class=\"font-bold \">" + fav.Title + "</span>				</div>				<p class=\"text-gray-700 text-xs\">" + fav.Category + "</p>			</div>		</div>		<p class=\"w-2/5\">" + fav.Overview + "</p>		<label for=\"stars\" class=\"font-bold w-1/5 text-right\">" + fav.Stars + "</label>	</li></a>\n"
	}
	fmt.Println("html:" + html)
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
