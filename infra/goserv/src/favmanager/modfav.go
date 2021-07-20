package favmanager

import (
	"database/sql"
	"errors"
	"myfav/crtuser"
	"myfav/sessionmanager"
	"myfav/utils"

	"github.com/gin-gonic/gin"
)

// Favmod Favの新規登録
func Favmod(c *gin.Context, fav Fav) error {
	var err error
	fav.Userid, err = sessionmanager.GetUserId(c)
	if err != nil {
		return errors.New("favadd getuserid:" + err.Error())
	}

	fav.Timing = timingconv(fav.Timing)
	fav.Stars = starsconv(fav.Stars)
	fav.Openclose = opclconv(fav.Openclose)

	if fav.Title == "" {
		return errors.New("favadd:title is empty")
	}

	if err := updateFav(fav); err != nil {
		return errors.New("favadd updateFav:" + err.Error())
	}
	return nil
}

func updateFav(fav Fav) error {
	db, err := sql.Open(utils.DBName, utils.ConnectStringDB)
	if err != nil {
		return errors.New("updatefav:sqlopen " + err.Error())
	}
	defer db.Close()

	if rs, err := updateFavs_favs(db, fav); err != nil {
		return err
	} else if cnt, err := rs.RowsAffected(); cnt != 0 {
		return updateFavs_image(db, fav, rs)
	} else if cnt == 0 {
		return errors.New("updateFav:更新レコードが0件です。")
	} else {
		return err
	}
}

func updateFavs_favs(db *sql.DB, fav Fav) (sql.Result, error) {
	stmtUpdate, err := db.Prepare(utils.FavUpdateSQL)
	if err != nil {
		return nil, errors.New("updateFavs_favs " + err.Error())
	}
	defer stmtUpdate.Close()

	var rs sql.Result
	rs, err = stmtUpdate.Exec(fav.Title, fav.Category, fav.Publisher, fav.Overview,
		fav.Impre, fav.Timing, fav.Stars, fav.Openclose, crtuser.GetTimeString(), fav.Userid, fav.Favid)
	return rs, err
}

func updateFavs_image(db *sql.DB, fav Fav, rs sql.Result) error {
	stmtInsert, err := db.Prepare(utils.ImageUpdateSQL)
	if err != nil {
		return errors.New("updateFavs_image " + err.Error())
	}
	defer stmtInsert.Close()
	_, err = stmtInsert.Exec(fav.Icon, fav.Favid)
	return err
}
