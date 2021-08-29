package identifychk

import (
	"crypto/sha256"
	"errors"

	"myfav/dbaccessor"

	_ "github.com/go-sql-driver/mysql"
)

func Invalidchk(username string, password string) error {
	hashpass := sha256.Sum256([]byte(password))
	dbpasshash, err := dbaccessor.GetUserpass(username)
	if err != nil {
		return err
	}
	for i := 0; i < len(dbpasshash); i++ {
		if hashpass[i] != dbpasshash[i] {
			return errors.New("ユーザーIDかパスワードが一致しません。")
		}
	}
	return nil
}
