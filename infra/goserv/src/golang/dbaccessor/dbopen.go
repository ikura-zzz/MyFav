package dbaccessor

import (
	"database/sql"
	"errors"
	"myfav/logmanager"
	"myfav/utils"
)

func DBOpen() (*sql.DB, error) {
	db, err := sql.Open(utils.DBName, utils.ConnectStringDB)
	if err != nil {
		logmanager.Outlog("sqlopen failed")
		return nil, errors.New("sql.open " + err.Error())
	}
	return db, nil
}
