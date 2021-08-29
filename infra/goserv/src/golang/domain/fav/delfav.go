package fav

import (
	"errors"

	"myfav/dbaccessor"
	"myfav/sessionmanager"
	"myfav/types"
	"myfav/utils"

	"github.com/gin-gonic/gin"
)

// Favdel Favの削除
func Favdel(c *gin.Context, fav types.Fav) error {
	var err error
	fav.Userid, err = sessionmanager.GetUserId(c)
	if err != nil {
		return errors.New("favdel getuserid:" + err.Error())
	}
	return deleteFav(fav)
}

func deleteFav(fav types.Fav) error {
	db, err := dbaccessor.DBOpen()
	if err != nil {
		return err
	}
	defer db.Close()

	favs, err := dbaccessor.Selectfavs(fav.Userid, utils.SelectFavsByUserid)
	if err != nil {
		return errors.New("deleteFav:selectFavs:" + err.Error())
	}
	for _, f := range favs {
		if f.Favid == fav.Favid {
			return dbaccessor.ExecDeleteFav(db, fav)
		}
	}
	return errors.New("deleteFav:削除対象レコードが0件です。")
}
