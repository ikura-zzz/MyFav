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

// ユーザー追加の実行
func Favadd(c *gin.Context, fav Fav) error {
	var err error
	fav.Userid, err = getusrid(c)
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

	return "0"
}
func starsconv(stars string) string {
	if stars == "one" {
		return "1"
	}
	if stars == "two" {
		return "2"
	}
	if stars == "three" {
		return "3"
	}
	if stars == "four" {
		return "4"
	}
	if stars == "five" {
		return "5"
	}

	return "0"
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
func getusrid(c *gin.Context) (int, error) {

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
		//return err
		return errors.New("resistfav1" + err.Error())
	}
	defer db.Close()

	stmtInsert, err := db.Prepare(utils.FavInsertSQL)
	if err != nil {
		//return err
		return errors.New("resistfav2" + err.Error())
	}
	defer stmtInsert.Close()

	_, err = stmtInsert.Exec(fav.Userid, fav.Title, fav.Category, fav.Publisher, fav.Overview, fav.Impre, fav.Timing, fav.Stars, fav.Openclose, crtusr.GetTimeString())
	return err

}
