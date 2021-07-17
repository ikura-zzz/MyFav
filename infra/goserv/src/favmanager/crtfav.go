package favmanager

import (
	"database/sql"
	"errors"
	"myfav/crtusr"
	"myfav/identifychk"
	"myfav/sessionmanager"
	"myfav/utils"

	"github.com/gin-gonic/gin"
)

// Favadd Favの新規登録
func Favadd(c *gin.Context, fav Fav) error {
	var err error
	fav.Userid, err = Getusrid(c)
	if err != nil {
		return err
	}

	fav.Timing = timingconv(fav.Timing)
	fav.Stars = starsconv(fav.Stars)
	fav.Openclose = opclconv(fav.Openclose)

	if err := resistFav(fav); err != nil {
		//return err
		return errors.New("favadd2" + err.Error())
	}
	return nil
}
func timingconv(timing string) string {
	if timing == "already" {
		return "1"
	}
	if timing == "now" {
		return "2"
	}
	if timing == "wish" {
		return "3"
	}

	return "2"
}

func starsconv(stars string) string {
	if stars == "1" {
		return "1"
	}
	if stars == "2" {
		return "2"
	}
	if stars == "3" {
		return "3"
	}
	if stars == "4" {
		return "4"
	}
	if stars == "5" {
		return "5"
	}

	return "3"
}
func opclconv(opcl string) string {
	if opcl == "open" {
		return "1"
	}
	if opcl == "close" {
		return "0"
	}

	return "0"
}
func Getusrid(c *gin.Context) (int, error) {

	username, ok := sessionmanager.GetSession(c, "username")
	if !ok {
		return 0, errors.New("DBアクセス不可")
	}
	id, err := identifychk.GetUsrId(username)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func resistFav(fav Fav) error {
	db, err := sql.Open(utils.DBName, utils.ConnectStringDB)
	if err != nil {
		return errors.New("resistfav:sqlopen " + err.Error())
	}
	defer db.Close()

	if rs, err := resistFavs_favs(db, fav); err != nil {
		return err
	} else {
		return resistFavs_image(db, fav, rs)
	}
}

func resistFavs_favs(db *sql.DB, fav Fav) (sql.Result, error) {
	stmtInsert, err := db.Prepare(utils.FavInsertSQL)
	if err != nil {
		return nil, errors.New("resistFavs_favs " + err.Error())
	}
	defer stmtInsert.Close()

	var rs sql.Result
	rs, err = stmtInsert.Exec(fav.Userid, fav.Title, fav.Category, fav.Publisher, fav.Overview, fav.Impre, fav.Timing, fav.Stars, fav.Openclose, crtusr.GetTimeString())
	return rs, err
}

func resistFavs_image(db *sql.DB, fav Fav, rs sql.Result) error {
	stmtInsert, err := db.Prepare(utils.ImageInsertSQL)
	if err != nil {
		return errors.New("resistFavs_image " + err.Error())
	}
	defer stmtInsert.Close()

	favid, _ := rs.LastInsertId()
	_, err = stmtInsert.Exec(favid, fav.Icon)
	return err
}