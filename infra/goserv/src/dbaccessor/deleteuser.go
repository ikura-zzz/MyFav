package dbaccessor

import (
	"database/sql"
	"errors"
	"myfav/logmanager"
	"myfav/utils"
)

func DeleteUser(db *sql.DB, userid int) error {
	stmtDelete, err := db.Prepare(utils.UserDeleteSQL)
	if err != nil {
		logmanager.Outlog("deleteUser:stmtcreate:" + err.Error())
		return errors.New("deleteUser:stmtcreate:" + err.Error())
	}
	defer stmtDelete.Close()

	_, err = stmtDelete.Exec(userid)
	if err != nil {
		logmanager.Outlog("delteUser:sqlExec:" + err.Error())
		return errors.New("delteUser:sqlExec:" + err.Error())
	}
	return nil
}
