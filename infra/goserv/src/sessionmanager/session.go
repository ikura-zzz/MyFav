package sessionmanager

import (
	"errors"

	"myfav/identifychk"
	"myfav/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const sessionname string = "myfavsession"
const sessionMaxAge int = 1800

func CreateNewSessionStore(engine *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))
	engine.Use(sessions.Sessions(sessionname, store))
}

func CreateNewSession(c *gin.Context, username string) {
	session := sessions.Default(c)
	session.Options(sessions.Options{MaxAge: sessionMaxAge})
	session.Set(utils.SessionKeyUser, username)
	session.Save()
}
func GetSessionValue(c *gin.Context, key string) (string, bool) {
	session := sessions.Default(c)
	session.Options(sessions.Options{MaxAge: sessionMaxAge})
	value, ok := session.Get(key).(string)
	return value, ok
}
func InqSessionValid(c *gin.Context, key string) bool {
	value, ok := GetSessionValue(c, key)
	if !ok {
		return false
	}
	if cnt, err := identifychk.GetUserCnt(value); err != nil {
		return false
	} else {
		return cnt > 0
	}
}

func RemoveSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}
func GetUserId(c *gin.Context) (int, error) {

	username, ok := GetSessionValue(c, "username")
	if !ok {
		return 0, errors.New("DBアクセス不可")
	}
	id, err := identifychk.GetUserId(username)
	if err != nil {
		return 0, err
	}
	return id, nil
}
