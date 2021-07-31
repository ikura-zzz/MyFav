package usermanager

import (
	"errors"
	"myfav/identifychk"
	"myfav/logmanager"
	"myfav/utils"
)

// Useradd ユーザー新規追加
func Useradd(username string, pass1 string, pass2 string) error {
	if err := passwordValid(pass1, pass2); err != nil {
		return err
	}
	if cnt, err := identifychk.GetUserCnt(username); err != nil {
		logmanager.Outlog(err.Error())
		return errors.New(utils.CmnErrmsg)
	} else if cnt != 0 {
		logmanager.Outlog("same username already used.")
		return errors.New("このユーザー名は既に使用されています。")
	}
	if err := registUser(username, pass1); err != nil {
		logmanager.Outlog(err.Error())
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
