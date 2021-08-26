package usermanager

import (
	"errors"
	"fmt"

	"myfav/dbaccessor"
	"myfav/domain/logger"
	"myfav/myfavtime"
	"myfav/utils"
)

func AppUsersMod(userid int, newname string) error {
	db, err := dbaccessor.DBOpen()
	if err != nil {
		return err
	}
	defer db.Close()

	stmtUpdate, err := db.Prepare(utils.UserNameUpdateSQL)
	if err != nil {
		return err
	}
	defer stmtUpdate.Close()

	rs, err := stmtUpdate.Exec(newname, myfavtime.GetTimeString(), userid)
	if err != nil {
		return errors.New("AppUsersMod:" + err.Error())
	}
	if cnt, _ := rs.RowsAffected(); cnt != 1 {
		return errors.New("更新件数障害 更新件数：" + fmt.Sprintf("%d", cnt))
	}
	return nil
}

func AppUserspassMod(userid int, newhashpass [32]byte) error {
	var l logger.Logger = new(logger.Logimp)
	db, err := dbaccessor.DBOpen()
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

	rs, err := stmtUpdate.Exec(fmt.Sprintf("%s", newhashpass), myfavtime.GetTimeString(), userid)
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
