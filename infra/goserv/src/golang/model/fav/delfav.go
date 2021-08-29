package fav

import (
	"errors"
	"myfav/dbaccessor"
	"myfav/types"
	"myfav/utils"
)

// Favdel Favの削除
func Favdel(fav types.Fav) error {
	favs, err := dbaccessor.Selectfavs(fav.Userid, utils.SelectFavsByUserid)
	if err != nil {
		return errors.New("deleteFav:selectFavs:" + err.Error())
	}
	for _, f := range favs {
		if f.Favid == fav.Favid {
			return dbaccessor.DeleteFav(fav)
		}
	}
	return errors.New("deleteFav:削除対象レコードが0件です。")
}
