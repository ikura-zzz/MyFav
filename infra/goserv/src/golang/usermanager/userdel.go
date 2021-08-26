package usermanager

import (
	"database/sql"
	"fmt"
	"myfav/dbaccessor"
	"myfav/logmanager"
	"myfav/types"
	"myfav/utils"
)

// Userdel ユーザー削除
func Userdel(userid int) error {
	logmanager.Outlog("userdel:userid:" + fmt.Sprintf("%d", userid))
	favs, err := dbaccessor.Selectfavs(userid, utils.SelectFavsByUserid)
	if err != nil {
		logmanager.Outlog("Userdel:selectfavs" + err.Error())
		return err
	}

	db, err := dbaccessor.DBOpen()
	if err != nil {
		return err
	}
	defer db.Close()

	logmanager.Outlog("fav cnt:" + fmt.Sprintf("%d", len(favs)))
	if len(favs) > 0 {
		if err := delfavs(db, userid, favs); err != nil {
			return err
		}
	}
	return dbaccessor.DeleteUser(db, userid)
}

func delfavs(db *sql.DB, userid int, favs []types.Fav) error {

	for _, f := range favs {
		f.Userid = userid
		if err := dbaccessor.ExecDeleteFav(db, f); err != nil {
			return err
		}
	}
	return nil
}
