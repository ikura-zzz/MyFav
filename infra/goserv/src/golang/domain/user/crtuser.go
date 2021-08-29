package user

import (
	"errors"
	"myfav/dbaccessor"
	"myfav/domain/logger"
	"myfav/utils"
)

// Useradd ユーザー新規追加
func Useradd(username string, pass1 string, pass2 string) error {
	var l logger.Logger = new(logger.Logimp)
	if err := passwordValid(pass1, pass2); err != nil {
		return err
	}
	if cnt, err := dbaccessor.GetUserCnt(username); err != nil {
		l.Outlog(err.Error())
		return errors.New(utils.CmnErrmsg)
	} else if cnt != 0 {
		l.Outlog("same username already used.")
		return errors.New("このユーザー名は既に使用されています。")
	}
	if err := registUser(username, pass1); err != nil {
		l.Outlog(err.Error())
		return errors.New(utils.CmnErrmsg)
	}
	return nil
}

// パスワード登録時の1つ目と2つ目のパスワード入力が一致しているか判定
func passwordValid(pass1 string, pass2 string) error {
	if pass1 != pass2 {
		return errors.New("パスワードが一致しません。")
	}
	return nil
}
