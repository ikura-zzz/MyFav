package crtusr

import (
	// mysql driver

	"crypto/sha256"
	"database/sql"
	"fmt"
	"myfav/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// registUser useradd to db
func registUser(username string, password string) error {
	hashpass := sha256.Sum256([]byte(password))
	db, err := sql.Open(utils.DBName, utils.ConnectStringDB)
	if err != nil {
		return err
	}
	defer db.Close()

	stmtInsert, err := db.Prepare(utils.UserInsertSQL)
	if err != nil {
		return err
	}
	defer stmtInsert.Close()

	_, err = stmtInsert.Exec(username, fmt.Sprintf("%s", hashpass), getTimeString())
	return err
}

func getTimeString() string {
	layout := "2006-01-02 15:04:05"
	now := time.Now()
	return now.Format(layout)
}
