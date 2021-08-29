package user

import (
	"crypto/sha256"
	"errors"

	"myfav/dbaccessor"
	"myfav/types"

	_ "github.com/go-sql-driver/mysql"
)

// Invalidchk 入力されたユーザー名とパスワードの一致を確認する。
func Invalidchk(username string, password string) error {
	hashpass := (types.Hashpass)(sha256.Sum256([]byte(password)))
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
