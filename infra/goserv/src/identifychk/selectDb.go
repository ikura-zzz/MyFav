package selectdb

import (
	"database/sql"
	"errors"
	"myfav/utils"
)

func Invalidchk(username string, inputpasshash [32]byte) error {
	db, err := sql.Open(utils.DBName, utils.ConnectStringDB)
	if err != nil {
		return err
	}
	defer db.Close()

	stmtInsert, err := db.Prepare(utils.SelectUserPass)
	if err != nil {
		return err
	}
	defer stmtInsert.Close()

	rows, err := stmtInsert.Query(username)
	if err != nil {
		return err
	}
	defer rows.Close()

	rows.Next()
	var dbpasshashstr string
	err = rows.Scan(&dbpasshashstr)
	if err != nil {
		return err
	}
	dbpasshash := []byte(dbpasshashstr)

	for i := 0; i < len(dbpasshash); i++ {
		if inputpasshash[i] != dbpasshash[i] {
			return errors.New("パスワード不一致")
		}
	}
	return nil
}

func GetUsrCnt(username string) (int, error) {
	db, err := sql.Open(utils.DBName, utils.ConnectStringDB)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	stmtInsert, err := db.Prepare(utils.SelectUsernameCnt)
	if err != nil {
		return 0, err
	}
	defer stmtInsert.Close()

	rows, err := stmtInsert.Query(username)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	rows.Next()
	var cnt int
	err = rows.Scan(&cnt)
	if err != nil {
		return 0, err
	}
	return cnt, nil
}
