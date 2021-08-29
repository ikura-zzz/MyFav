package user

import (
	"fmt"
	"myfav/dbaccessor"
	"myfav/domain/logger"
	"myfav/utils"
)

// Userdel ユーザー削除
func Userdel(userid int) error {
	var l logger.Logger = new(logger.Logimp)
	l.Outlog("userdel:userid:" + fmt.Sprintf("%d", userid))
	favs, err := dbaccessor.Selectfavs(userid, utils.SelectFavsByUserid)
	if err != nil {
		l.Outlog("Userdel:selectfavs" + err.Error())
		return err
	}

	l.Outlog("fav cnt:" + fmt.Sprintf("%d", len(favs)))
	if len(favs) > 0 {
		for _, f := range favs {
			f.Userid = userid
			if err := dbaccessor.DeleteFav(f); err != nil {
				return err
			}
		}
	}
	return dbaccessor.DeleteUser(userid)
}
