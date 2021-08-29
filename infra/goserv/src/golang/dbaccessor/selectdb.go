package dbaccessor

import (
	"errors"
	"myfav/types"
	"myfav/utils"
)

// Selectfavs FavをDBから取得する。
func Selectfavs(userid int, selectSql string) ([]types.Fav, error) {
	db, err := dbOpen()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmtInsert, err := db.Prepare(selectSql)
	if err != nil {
		return nil, errors.New("db.Prepare " + err.Error())
	}
	defer stmtInsert.Close()

	rows, err := stmtInsert.Query(userid)
	if err != nil {
		return nil, errors.New("stmt.Query " + err.Error())
	}
	defer rows.Close()

	var favs []types.Fav

	for i := 0; rows.Next(); i++ {
		fav := types.Fav{}
		err := rows.Scan(&fav.Favid, &fav.Title, &fav.Category, &fav.Publisher, &fav.Overview,
			&fav.Impre, &fav.Timing, &fav.Stars, &fav.Openclose, &fav.Icon)
		if err != nil {
			return nil, errors.New("SelectfavsByUserid gerfavs:" + err.Error())
		}
		favs = append(favs, fav)
	}
	return favs, nil
}

func GetUserpass(username string) ([]byte, error) {
	db, err := dbOpen()
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
	db, err := dbOpen()
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
	db, err := dbOpen()
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
