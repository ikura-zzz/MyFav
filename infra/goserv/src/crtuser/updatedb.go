package crtuser

import (
	"database/sql"
	"errors"
	"fmt"
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
