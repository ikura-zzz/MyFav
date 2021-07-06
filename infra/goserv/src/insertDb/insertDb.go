package insertdb

import (
	// mysql driver

	"os/exec"

	_ "github.com/go-sql-driver/mysql"
)

// InsertDb return error
func InsertDb(inputs []string) ([]byte, error) {
	cmd := exec.Command("adddb", "-s", inputs[0], inputs[1], inputs[2],
		inputs[3], inputs[4], inputs[5])
	output, err := cmd.CombinedOutput()
	if err != nil {
		return output, err
	}

	return nil, nil
}

// RegistUser add user
func RegistUser(username string, password string) ([]byte, error) {
	cmd := exec.Command("adddb", "-u", username, password)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return output, err
	}
	return nil, nil
}
