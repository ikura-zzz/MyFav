package usermanager

import (
	// mysql driver

	"crypto/sha256"
	"fmt"

	"myfav/dbaccessor"
	"myfav/myfavtime"
	"myfav/utils"

	_ "github.com/go-sql-driver/mysql"
)

// registUser useradd to db
func registUser(username string, password string) error {
	hashpass := sha256.Sum256([]byte(password))
	db, err := dbaccessor.DBOpen()
	if err != nil {
		return err
	}
	defer db.Close()

	stmtInsert, err := db.Prepare(utils.UserInsertSQL)
	if err != nil {
		return err
	}
	defer stmtInsert.Close()

	_, err = stmtInsert.Exec(username, fmt.Sprintf("%s", hashpass), myfavtime.GetTimeString())
	return err
}
