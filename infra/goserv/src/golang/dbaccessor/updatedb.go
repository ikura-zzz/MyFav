package dbaccessor

import (
	"errors"
	"fmt"

	"myfav/model/logger"
	"myfav/types"
	"myfav/utils"
)

func AppUsersMod(userid int, newname string) error {
	db, err := dbOpen()
	if err != nil {
		return err
	}
	defer db.Close()

	stmtUpdate, err := db.Prepare(utils.UserNameUpdateSQL)
	if err != nil {
		return err
	}
	defer stmtUpdate.Close()

	rs, err := stmtUpdate.Exec(newname, utils.GetTimeString(), userid)
	if err != nil {
		return errors.New("AppUsersMod:" + err.Error())
	}
	if cnt, _ := rs.RowsAffected(); cnt != 1 {
		return errors.New("更新件数障害 更新件数：" + fmt.Sprintf("%d", cnt))
	}
	return nil
}

func AppUserspassMod(userid int, newhashpass types.Hashpass) error {
	var l logger.Logger = new(logger.Logimp)
	db, err := dbOpen()
	if err != nil {
		return err
	}
	defer db.Close()

	stmtUpdate, err := db.Prepare(utils.UserPasswordUpdateSQL)
	if err != nil {
		l.Outlog("AppUserspassmod:stmtPrepare:" + err.Error())
		return err
	}
	defer stmtUpdate.Close()

	rs, err := stmtUpdate.Exec(fmt.Sprintf("%s", newhashpass), utils.GetTimeString(), userid)
	if err != nil {
		l.Outlog("AppUserspassmod:stmtExec:" + err.Error())
		return errors.New("AppUserspassMod:" + err.Error())
	}
	if cnt, _ := rs.RowsAffected(); cnt != 1 {
		l.Outlog("AppUserspassMod:updated is:" + fmt.Sprintf("%d", cnt))
		return errors.New("更新件数障害 更新件数：" + fmt.Sprintf("%d", cnt))
	}
	return nil
}
