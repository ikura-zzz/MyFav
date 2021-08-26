package dbaccessor

import (
	"database/sql"
	"errors"
	"myfav/domain/logger"
	"myfav/utils"
)

func DeleteUser(db *sql.DB, userid int) error {
	stmtDelete, err := db.Prepare(utils.UserDeleteSQL)
	var l logger.Logger = new(logger.Logimp)
	if err != nil {
		l.Outlog("deleteUser:stmtcreate:" + err.Error())
		return errors.New("deleteUser:stmtcreate:" + err.Error())
	}
	defer stmtDelete.Close()

	_, err = stmtDelete.Exec(userid)
	if err != nil {
		l.Outlog("delteUser:sqlExec:" + err.Error())
		return errors.New("delteUser:sqlExec:" + err.Error())
	}
	return nil
}
