package dbaccessor

import (
	"errors"
	"myfav/model/logger"
	"myfav/utils"
)

// DeleteUser ユーザーを1件削除する。
func DeleteUser(userid int) error {
	db, err := dbOpen()
	if err != nil {
		return err
	}
	defer db.Close()
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
