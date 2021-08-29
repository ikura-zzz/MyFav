package dbaccessor

import (
	"crypto/sha256"
	"fmt"
	"myfav/utils"
)

// RegistUser useradd to db
func RegistUser(username string, password string) error {
	hashpass := sha256.Sum256([]byte(password))
	db, err := dbOpen()
	if err != nil {
		return err
	}
	defer db.Close()

	stmtInsert, err := db.Prepare(utils.UserInsertSQL)
	if err != nil {
		return err
	}
	defer stmtInsert.Close()

	_, err = stmtInsert.Exec(username, fmt.Sprintf("%s", hashpass), utils.GetTimeString())
	return err
}
