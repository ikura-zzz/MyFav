package session

import (
	"errors"

	"myfav/dbaccessor"
	"myfav/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const sessionname string = "myfavsession"
const sessionMaxAge int = 1800

// CreateNewSessionStore 新しいセッションストアを作る。
func CreateNewSessionStore(engine *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))
	engine.Use(sessions.Sessions(sessionname, store))
}

// CreateNewSession 新しいセッションを作る。
func CreateNewSession(c *gin.Context, username string) {
	session := sessions.Default(c)
	session.Options(sessions.Options{MaxAge: sessionMaxAge})
	session.Set(utils.SessionKeyUser, username)
	session.Save()
}

// GetSessionValue 与えられたkey値でセッションバリューを取り出して返す。
// 与えられたキーで値が取得できなかった場合は2つ目の戻り値でfalseを返す。
func GetSessionValue(c *gin.Context, key string) (string, bool) {
	session := sessions.Default(c)
	session.Options(sessions.Options{MaxAge: sessionMaxAge})
	value, ok := session.Get(key).(string)
	return value, ok
}

// InqSessionValid 与えられたコンテキストから、セッションが有効か判定して返す。
func InqSessionValid(c *gin.Context) bool {
	value, ok := GetSessionValue(c, utils.SessionKeyUser)
	if !ok {
		return false
	}
	if cnt, err := dbaccessor.GetUserCnt(value); err != nil {
		return false
	} else {
		return cnt > 0
	}
}

// RemoveSession セッションを削除する。
func RemoveSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

// GetUserId セッションからユーザーIDを取得して返す。
func GetUserId(c *gin.Context) (int, error) {
	username, ok := GetSessionValue(c, "username")
	if !ok {
		return 0, errors.New("DBアクセス不可")
	}
	id, err := dbaccessor.GetUserId(username)
	if err != nil {
		return 0, err
	}
	return id, nil
}
