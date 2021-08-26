package dbaccessor

import (
	"database/sql"
	"errors"
	"myfav/domain/logger"
	"myfav/utils"
)

func DBOpen() (*sql.DB, error) {
	db, err := sql.Open(utils.DBName, utils.ConnectStringDB)
	if err != nil {
		var l logger.Logger = new(logger.Logimp)
		l.Outlog("sqlopen failed")
		return nil, errors.New("sql.open " + err.Error())
	}
	return db, nil
}
