package user

import (
	"crypto/sha256"
	"errors"

	"myfav/dbaccessor"
	"myfav/domain/logger"
	"myfav/utils"
)

// Usernamemod ユーザー名変更
func Usernamemod(userid int, currentusername string, newusername string) error {
	var l logger.Logger = new(logger.Logimp)
	if currentusername == newusername {
		l.Outlog("this username is using now.")
		return errors.New("現在と同じユーザー名です。")
	}
	if cnt, err := dbaccessor.GetUserCnt(newusername); err != nil {
		l.Outlog(err.Error())
		return errors.New(utils.CmnErrmsg)
	} else if cnt != 0 {
		l.Outlog("same username already used.")
		return errors.New("このユーザー名は既に使用されています。")
	}
	if err := AppUsersMod(userid, newusername); err != nil {
		l.Outlog(("usernamemod:") + err.Error())
		return errors.New(utils.CmnErrmsg)
	}
	return nil
}

// Userpassmod ユーザーパスワード変更
func Userpassmod(userid int, username string, newpass string, retypepass string) error {
	var l logger.Logger = new(logger.Logimp)
	if err := passwordValid(newpass, retypepass); err != nil {
		l.Outlog("Userpassmod" + err.Error())
		return err
	}
	hashpass, err := dbaccessor.GetUserpass(username)
	var currenthashpass [32]byte
	for i := 0; i < len(currenthashpass); i++ {
		currenthashpass[i] = hashpass[i]
	}
	newhashpass := sha256.Sum256([]byte(newpass))
	if err != nil {
		l.Outlog("Userpassmod:" + err.Error())
		return errors.New(utils.CmnErrmsg)
	}
	if currenthashpass == newhashpass {
		l.Outlog("this password is using now.")
		return errors.New("このパスワードは現在使用中のものです。")
	}
	if err := AppUserspassMod(userid, newhashpass); err != nil {
		return errors.New(utils.CmnErrmsg)
	}
	return nil
}
