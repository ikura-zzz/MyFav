package usermanager

import (
	"database/sql"
	"errors"
	"fmt"
	"myfav/logmanager"
	"myfav/utils"
)

func AppUsersMod(userid int, newname string) error {
	db, err := sql.Open(utils.DBName, utils.ConnectStringDB)
	if err != nil {
		return err
	}
	defer db.Close()

	stmtUpdate, err := db.Prepare(utils.UserNameUpdateSQL)
	if err != nil {
		return err
	}
	defer stmtUpdate.Close()

	rs, err := stmtUpdate.Exec(newname, GetTimeString(), userid)
	if err != nil {
		return errors.New("AppUsersMod:" + err.Error())
	}
	if cnt, _ := rs.RowsAffected(); cnt != 1 {
		return errors.New("更新件数障害 更新件数：" + fmt.Sprintf("%d", cnt))
	}
	return nil
}

func AppUserspassMod(userid int, newhashpass [32]byte) error {
	db, err := sql.Open(utils.DBName, utils.ConnectStringDB)
	if err != nil {
		logmanager.Outlog("AppUserspassmod:sqlOpen:" + err.Error())
		return err
	}
	defer db.Close()

	stmtUpdate, err := db.Prepare(utils.UserPasswordUpdateSQL)
	if err != nil {
		logmanager.Outlog("AppUserspassmod:stmtPrepare:" + err.Error())
		return err
	}
	defer stmtUpdate.Close()

	rs, err := stmtUpdate.Exec(fmt.Sprintf("%s", newhashpass), GetTimeString(), userid)
	if err != nil {
		logmanager.Outlog("AppUserspassmod:stmtExec:" + err.Error())
		return errors.New("AppUserspassMod:" + err.Error())
	}
	if cnt, _ := rs.RowsAffected(); cnt != 1 {
		logmanager.Outlog("AppUserspassMod:updated is:" + fmt.Sprintf("%d", cnt))
		return errors.New("更新件数障害 更新件数：" + fmt.Sprintf("%d", cnt))
	}
	return nil
}
