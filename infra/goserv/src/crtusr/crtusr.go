package crtusr

import (
	"errors"
	selectdb "myfav/selectDb"
)

// ユーザー追加の実行
func Useradd(username string, pass1 string, pass2 string) error {
	cmnerrmsg := "ユーザー登録時に予期せぬエラーが発生しました。"
	if err := passwordValid(pass1, pass2); err != nil {
		return err
	}
	if cnt, err := selectdb.GetUsrCnt(username); err != nil {
		return errors.New(cmnerrmsg)
	} else if cnt != 0 {
		return errors.New("このユーザー名は既に使用されています。")
	}
	if err := registUser(username, pass1); err != nil {
		return errors.New(cmnerrmsg)
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
