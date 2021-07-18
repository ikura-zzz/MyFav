package favmanager

import (
	"database/sql"
	"errors"
	"myfav/utils"
)

func SelectfavsByUserid(userid int, selectSql string) ([]Fav, error) {
	db, err := sql.Open(utils.DBName, utils.ConnectStringDB)
	if err != nil {
		return nil, errors.New("sql.open " + err.Error())
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

	var favs []Fav

	for i := 0; rows.Next(); i++ {
		fav := Fav{}
		err := rows.Scan(&fav.Favid, &fav.Title, &fav.Category, &fav.Publisher, &fav.Overview,
			&fav.Impre, &fav.Timing, &fav.Stars, &fav.Openclose, &fav.Icon)
		if err != nil {
			return nil, errors.New("SelectfavsByUserid gerfavs:" + err.Error())
		}
		favs = append(favs, fav)
	}
	return favs, nil
}
