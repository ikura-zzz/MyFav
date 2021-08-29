package dbaccessor

import (
	"database/sql"
	"errors"
	"myfav/types"
	"myfav/utils"
)

func UpdateFav(fav types.Fav) error {
	db, err := dbOpen()
	if err != nil {
		return err
	}
	defer db.Close()

	if rs, err := updateFav_fav(db, fav); err != nil {
		return err
	} else if cnt, err := rs.RowsAffected(); cnt != 0 {
		return updateFav_image(db, fav)
	} else if cnt == 0 {
		return errors.New("updateFav:更新レコードが0件です。")
	} else {
		return err
	}
}

func updateFav_fav(db *sql.DB, fav types.Fav) (sql.Result, error) {
	stmtUpdate, err := db.Prepare(utils.FavUpdateSQL)
	if err != nil {
		return nil, errors.New("updateFav_fav:" + err.Error())
	}
	defer stmtUpdate.Close()

	var rs sql.Result
	rs, err = stmtUpdate.Exec(fav.Title, fav.Category, fav.Publisher, fav.Overview,
		fav.Impre, fav.Timing, fav.Stars, fav.Openclose, utils.GetTimeString(), fav.Userid, fav.Favid)
	return rs, err
}

func updateFav_image(db *sql.DB, fav types.Fav) error {
	stmtUpdate, err := db.Prepare(utils.ImageUpdateSQL)
	if err != nil {
		return errors.New("updateFav_image:" + err.Error())
	}
	defer stmtUpdate.Close()
	_, err = stmtUpdate.Exec(fav.Icon, fav.Favid)
	return err
}
