package dbaccessor

import (
	"errors"
	"myfav/types"
)

// Selectfavs FavをDBから取得する。
func Selectfavs(userid int, selectSql string) ([]types.Fav, error) {
	db, err := DBOpen()
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
