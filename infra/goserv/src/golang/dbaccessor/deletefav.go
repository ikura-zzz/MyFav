package dbaccessor

import (
	"database/sql"
	"errors"
	"myfav/types"
	"myfav/utils"
)

func Delfavs(userid int, favs []types.Fav) error {
	for _, f := range favs {
		f.Userid = userid
		if err := execDeleteFav(f); err != nil {
			return err
		}
	}
	return nil
}

func DeleteFav(fav types.Fav) error {

	favs, err := Selectfavs(fav.Userid, utils.SelectFavsByUserid)
	if err != nil {
		return errors.New("deleteFav:selectFavs:" + err.Error())
	}
	for _, f := range favs {
		if f.Favid == fav.Favid {
			return execDeleteFav(fav)
		}
	}
	return errors.New("deleteFav:削除対象レコードが0件です。")
}

func execDeleteFav(fav types.Fav) error {
	db, err := dbOpen()
	if err != nil {
		return err
	}
	defer db.Close()

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

func deleteFav_fav(db *sql.DB, fav types.Fav) error {
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

func deleteFav_image(db *sql.DB, fav types.Fav) (sql.Result, error) {
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
