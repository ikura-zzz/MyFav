package usermanager

import (
	"crypto/sha256"
	"errors"
	"myfav/identifychk"
	"myfav/logmanager"
	"myfav/utils"
)

// Usernamemod ユーザー名変更
func Usernamemod(userid int, currentusername string, newusername string) error {
	if currentusername == newusername {
		logmanager.Outlog("this username is using now.")
		return errors.New("現在と同じユーザー名です。")
	}
	if cnt, err := identifychk.GetUserCnt(newusername); err != nil {
		logmanager.Outlog(err.Error())
		return errors.New(utils.CmnErrmsg)
	} else if cnt != 0 {
		logmanager.Outlog("same username already used.")
		return errors.New("このユーザー名は既に使用されています。")
	}
	if err := AppUsersMod(userid, newusername); err != nil {
		logmanager.Outlog(("usernamemod:") + err.Error())
		return errors.New(utils.CmnErrmsg)
	}
	return nil
}

func Userpassmod(userid int, username string, newpass string, retypepass string) error {
	if err := passwordValid(newpass, retypepass); err != nil {
		logmanager.Outlog("Userpassmod" + err.Error())
		return errors.New(utils.CmnErrmsg)
	}
	hashpass, err := identifychk.GetUserpass(username)
	var currenthashpass [32]byte
	for i := 0; i < len(currenthashpass); i++ {
		currenthashpass[i] = hashpass[i]
	}
	newhashpass := sha256.Sum256([]byte(newpass))
	if err != nil {
		logmanager.Outlog("Userpassmod:" + err.Error())
		return errors.New(utils.CmnErrmsg)
	}
	if currenthashpass == newhashpass {
		logmanager.Outlog("this password is using now.")
		return errors.New("このパスワードは現在使用中のものです。")
	}
	if err := AppUserspassMod(userid, newhashpass); err != nil {
		return errors.New(utils.CmnErrmsg)
	}
	return nil
}
