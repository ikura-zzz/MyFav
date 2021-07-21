package favmanager

import (
	"database/sql"
	"errors"
	"myfav/sessionmanager"
	"myfav/utils"

	"github.com/gin-gonic/gin"
)

// Favdel Favの削除
func Favdel(c *gin.Context, fav Fav) error {
	var err error
	fav.Userid, err = sessionmanager.GetUserId(c)
	if err != nil {
		return errors.New("favdel getuserid:" + err.Error())
	}
	return deleteFav(fav)
}

func deleteFav(fav Fav) error {
	db, err := sql.Open(utils.DBName, utils.ConnectStringDB)
	if err != nil {
		return errors.New("deleteFav:sqlopen:" + err.Error())
	}
	defer db.Close()

	favs, err := Selectfavs(fav.Userid, utils.SelectFavsByUserid)
	if err != nil {
		return errors.New("deleteFav:selectFavs:" + err.Error())
	}
	for _, f := range favs {
		if f.Favid == fav.Favid {
			return execDeleteFav(db, fav)
		}
	}
	return errors.New("deleteFav:削除対象レコードが0件です。")
}
func execDeleteFav(db *sql.DB, fav Fav) error {
	if rs, err := deleteFav_image(db, fav); err != nil {
		return err
	} else if cnt, err := rs.RowsAffected(); cnt != 0 {
		return deleteFav_fav(db, fav)
	} else if cnt == 0 {
		return errors.New("deleteFav:削除レコードが0件です。")
	} else {
		return err
	}
}

func deleteFav_fav(db *sql.DB, fav Fav) error {
	stmtDelete, err := db.Prepare(utils.FavDeleteSQL)
	if err != nil {
		return errors.New("deleteFav_fav:opendb:" + err.Error())
	}
	defer stmtDelete.Close()

	_, err = stmtDelete.Exec(fav.Userid, fav.Favid)
	if err != nil {
		return errors.New("delteFav_fav:deletesql:" + err.Error())
	}
	return nil
}

func deleteFav_image(db *sql.DB, fav Fav) (sql.Result, error) {
	stmtDelete, err := db.Prepare(utils.ImageDeleteSQL)
	if err != nil {
		return nil, errors.New("deletefav_image:opendb:" + err.Error())
	}
	defer stmtDelete.Close()

	rs, err := stmtDelete.Exec(fav.Favid)
	if err != nil {
		return nil, errors.New("deleteFav_image:deletesql:" + err.Error())
	}
	return rs, nil
}
