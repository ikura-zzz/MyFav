package favmanager

import (
	"database/sql"
	"errors"

	"myfav/dbaccessor"
	"myfav/myfavtime"
	"myfav/sessionmanager"
	"myfav/types"
	"myfav/utils"

	"github.com/gin-gonic/gin"
)

// Favmod Favの編集
func Favmod(c *gin.Context, fav types.Fav) error {
	var err error
	fav.Userid, err = sessionmanager.GetUserId(c)
	if err != nil {
		return errors.New("favmod getuserid:" + err.Error())
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

func updateFav(fav types.Fav) error {
	db, err := dbaccessor.DBOpen()
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
		fav.Impre, fav.Timing, fav.Stars, fav.Openclose, myfavtime.GetTimeString(), fav.Userid, fav.Favid)
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
