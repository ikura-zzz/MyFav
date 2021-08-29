package fav

import (
	"errors"

	"myfav/dbaccessor"
	"myfav/domain/session"
	"myfav/types"

	"github.com/gin-gonic/gin"
)

// Favdel Favの削除
func Favdel(c *gin.Context, fav types.Fav) error {
	var err error
	fav.Userid, err = session.GetUserId(c)
	if err != nil {
		return errors.New("favdel getuserid:" + err.Error())
	}
	return dbaccessor.DeleteFav(fav)
}
