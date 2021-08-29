package dbaccessor

import (
	"database/sql"
	"errors"
	"myfav/model/logger"
)

func dbOpen() (*sql.DB, error) {
	db, err := sql.Open(dbName, connectStringDB)
	if err != nil {
		var l logger.Logger = new(logger.Logimp)
		l.Outlog("sqlopen failed")
		return nil, errors.New("sql.open " + err.Error())
	}
	return db, nil
}
