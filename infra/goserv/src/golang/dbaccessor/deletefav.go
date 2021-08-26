package dbaccessor

import (
	"database/sql"
	"errors"
	"myfav/types"
	"myfav/utils"
)

func ExecDeleteFav(db *sql.DB, fav types.Fav) error {
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
