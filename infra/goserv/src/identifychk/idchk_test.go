package identifychk

import (
	"database/sql"
	"myfav/utils"
	"testing"
)

func TestDBConn(t *testing.T) {
	db, err := sql.Open(utils.DBName, utils.ConnectStringDB)
	if err != nil {
		t.Fatalf("failed test%#v", err.Error())
	}
	defer db.Close()
}

func TestStmtInsert(t *testing.T) {
	db, err := sql.Open(utils.DBName, utils.ConnectStringDB)
	if err != nil {
		t.Fatalf("failed test%#v", err.Error())
	}
	defer db.Close()
	stmtInsert, err := db.Prepare(utils.SelectUserPass)
	if err != nil {
		t.Fatalf("failed test%#v", err.Error())
	}
	defer stmtInsert.Close()
}

func TestExecQuery(t *testing.T) {
	db, err := sql.Open(utils.DBName, utils.ConnectStringDB)
	if err != nil {
		t.Fatalf("failed test%#v", err.Error())
	}
	defer db.Close()

	stmtInsert, err := db.Prepare(utils.SelectUserPass)
	if err != nil {
		t.Fatalf("failed test%#v", err.Error())
	}
	defer stmtInsert.Close()

	rows, err := stmtInsert.Query("shigeji")
	if err != nil {
		t.Fatalf("failed test%#v", err.Error())
	}
	defer rows.Close()
}
