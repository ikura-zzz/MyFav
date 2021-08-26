package identifychk

import (
	"crypto/sha256"
	"errors"

	"myfav/dbaccessor"
	"myfav/utils"

	_ "github.com/go-sql-driver/mysql"
)

func Invalidchk(username string, password string) error {
	hashpass := sha256.Sum256([]byte(password))
	dbpasshash, err := GetUserpass(username)
	if err != nil {
		return err
	}
	for i := 0; i < len(dbpasshash); i++ {
		if hashpass[i] != dbpasshash[i] {
			return errors.New("ユーザーIDかパスワードが一致しません。")
		}
	}
	return nil
}
func GetUserpass(username string) ([]byte, error) {
	db, err := dbaccessor.DBOpen()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmtSelect, err := db.Prepare(utils.SelectUserPass)
	if err != nil {
		return nil, errors.New(utils.CmnErrmsg)
	}
	defer stmtSelect.Close()

	rows, err := stmtSelect.Query(username)
	if err != nil {
		return nil, errors.New(utils.CmnErrmsg)
	}
	defer rows.Close()

	rows.Next()
	var dbpasshashstr string
	err = rows.Scan(&dbpasshashstr)
	if err != nil {
		return nil, errors.New("このユーザーは存在しません。")
	}
	passhash := []byte(dbpasshashstr)
	return passhash, nil
}

func GetUserCnt(username string) (int, error) {
	db, err := dbaccessor.DBOpen()
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

func GetUserId(username string) (int, error) {
	db, err := dbaccessor.DBOpen()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	stmtInsert, err := db.Prepare(utils.SelectUserID)
	if err != nil {
		return 0, errors.New("db.Prepare " + err.Error())
	}
	defer stmtInsert.Close()

	rows, err := stmtInsert.Query(username)
	if err != nil {
		return 0, errors.New("stmt.Query " + err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		return 0, errors.New("not exist user")
	}
	var id int
	err = rows.Scan(&id)
	if err != nil {
		return 0, errors.New("row.Scan " + err.Error())
	}
	return id, nil
}