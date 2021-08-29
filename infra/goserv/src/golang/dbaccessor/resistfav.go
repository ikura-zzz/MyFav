package dbaccessor

import (
	"database/sql"
	"errors"
	"myfav/types"
	"myfav/utils"
)

func resistFavs_favs(db *sql.DB, fav types.Fav) (sql.Result, error) {
	stmtInsert, err := db.Prepare(utils.FavInsertSQL)
	if err != nil {
		return nil, errors.New("resistFavs_favs " + err.Error())
	}
	defer stmtInsert.Close()

	var rs sql.Result
	rs, err = stmtInsert.Exec(fav.Userid, fav.Title, fav.Category, fav.Publisher, fav.Overview, fav.Impre, fav.Timing, fav.Stars, fav.Openclose, utils.GetTimeString())
	return rs, err
}

func resistFavs_image(db *sql.DB, fav types.Fav, rs sql.Result) error {
	stmtInsert, err := db.Prepare(utils.ImageInsertSQL)
	if err != nil {
		return errors.New("resistFavs_image " + err.Error())
	}
	defer stmtInsert.Close()

	favid, _ := rs.LastInsertId()
	_, err = stmtInsert.Exec(favid, fav.Icon)
	return err
}

func ResistFav(fav types.Fav) error {
	db, err := DBOpen()
	if err != nil {
		return err
	}
	defer db.Close()

	if rs, err := resistFavs_favs(db, fav); err != nil {
		return err
	} else {
		return resistFavs_image(db, fav, rs)
	}
}
